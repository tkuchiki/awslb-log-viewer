package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	file           = kingpin.Flag("file", "logfile").Short('f').String()
	lbType         = kingpin.Flag("lb-type", "ALB or ELB").Short('t').Default("ALB").String()
	reALB          = regexp.MustCompile(`(.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) "(.+)" "(.+)" (.+) (.+) (.+) "(.+)" "(.+)" "(.+)" (.+) (.+) "(.+)" "(.+)" "(.+)" "(.+)" "(.+)"`)
	reRequest      = regexp.MustCompile(`(.+) (.+) (.+)`)
	reELB          = regexp.MustCompile(`(.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) "(.+)" "(.+)" (.+) (.+)`)
	invalidDataErr = errors.New("Invalid data")
)

func parseRequest(data string) (string, string, string, string, string, string, error) {
	req := reRequest.FindStringSubmatch(data)
	u, err := url.Parse(req[2])
	if err != nil {
		return "", "", "", "", "", "", err
	}

	hostPort := strings.Split(u.Host, `:`)

	return req[1], req[3], u.Scheme, hostPort[0], hostPort[1], u.Path, nil
}

func toJson(rawdata string, lbType string) (string, error) {
	var line string
	switch lbType {
	case "ALB":
		data := reALB.FindStringSubmatch(rawdata)
		if len(data) < 25 {
			return "", invalidDataErr
		}
		for i, _ := range data {
			data[i] = strings.Trim(data[i], `"`)
		}
		method, version, protocol, host, port, uri, err := parseRequest(data[13])
		if err != nil {
			return "", err
		}

		line = fmt.Sprintf(`{"type": "%s", "timestamp": "%s", "elb": "%s", "client": "%s", "target": "%s", "request_processing_time": "%s", "target_processing_time": "%s", "response_processing_time": "%s", "elb_status_code": "%s", "target_status_code": "%s", "received_bytes": "%s", "sent_bytes": "%s", "method": "%s", "http_version": "%s", "protocol": "%s", "host": "%s", "port": "%s", "uri": "%s", "user_agent": "%s", "ssl_cipher": "%s", "ssl_protocol": "%s", "target_group_arn": "%s", "trace_id": "%s", "domain_name": "%s", "chosen_cert_arn": "%s", "matched_rule_priority": "%s", "request_creation_time": "%s", "actions_executed": "%s", "redirect_url": "%s", "error_reason": "%s", "target:port_list": "%s", "target_status_code_list": "%s"}`,
			data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8], data[9], data[10], data[11], data[12], method, version, protocol, host, port, uri, data[14], data[15], data[16], data[17], data[18], data[19], data[20], data[21], data[22], data[23], data[24], data[25], data[26], data[27],
		)
	case "ELB":
		data := reELB.FindStringSubmatch(rawdata)
		if len(data) < 16 {
			return "", invalidDataErr
		}
		for i, _ := range data {
			data[i] = strings.Trim(data[i], `"`)
		}

		line = fmt.Sprintf(`{"timestamp": "%s", "elb": "%s", "client": "%s", "target": "%s", "request_processing_time": "%s", "backend_processing_time": "%s", "response_processing_time": "%s", "elb_status_code": "%s", "backend_status_code": "%s", "received_bytes": "%s", "sent_bytes": "%s", "reqeust": "%s", "user_agent": "%s", "ssl_cipher": "%s", "ssl_protocol": "%s"}`,
			data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8], data[9], data[10], data[11], data[12], data[13], data[14], data[15],
		)
	}

	return line, nil
}

func main() {
	kingpin.Version("0.1.1")
	kingpin.Parse()

	var f *os.File
	var err error
	stdinfi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if stdinfi.Mode()&os.ModeNamedPipe == 0 {
		f, err = os.Open(*file)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		f = os.Stdin
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line, lerr := toJson(scanner.Text(), *lbType)
		if lerr != nil {
			log.Fatal(lerr)
		}

		fmt.Println(line)
	}
}
