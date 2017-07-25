package main

import (
	"bufio"
	"errors"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	"regexp"
)

var (
	file           = kingpin.Flag("file", "logfile").Short('f').String()
	lbType         = kingpin.Flag("lb-type", "ALB or ELB").Short('t').Default("ALB").String()
	reALB          = regexp.MustCompile(`(.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) "(.+)" "(.+)" (.+) (.+) (.+) "(.+)"`)
	reELB          = regexp.MustCompile(`(.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) "(.+)" "(.+)" (.+) (.+)`)
	invalidDataErr = errors.New("Invalid data")
)

func toJson(rawdata string, lbType string) (string, error) {
	var line string
	switch lbType {
	case "ALB":
		data := reALB.FindStringSubmatch(rawdata)
		if len(data) < 19 {
			return "", invalidDataErr
		}
		line = fmt.Sprintf(`{"type": "%s", "timestamp": "%s", "elb": "%s", "client": "%s", "target": "%s", "request_processing_time": "%s", "target_processing_time": "%s", "response_processing_time": "%s", "elb_status_code": "%s", "target_status_code": "%s", "received_bytes": "%s", "sent_bytes": "%s", "reqeust": "%s", "user_agent": "%s", "ssl_cipher": "%s", "ssl_protocol": "%s", "target_group_arn": "%s", "trace_id": "%s"}`,
			data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8], data[9], data[10], data[11], data[12], data[13], data[14], data[15], data[16], data[17], data[18],
		)
	case "ELB":
		data := reELB.FindStringSubmatch(rawdata)
		if len(data) < 16 {
			return "", invalidDataErr
		}
		line = fmt.Sprintf(`{"timestamp": "%s", "elb": "%s", "client": "%s", "target": "%s", "request_processing_time": "%s", "backend_processing_time": "%s", "response_processing_time": "%s", "elb_status_code": "%s", "backend_status_code": "%s", "received_bytes": "%s", "sent_bytes": "%s", "reqeust": "%s", "user_agent": "%s", "ssl_cipher": "%s", "ssl_protocol": "%s"}`,
			data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8], data[9], data[10], data[11], data[12], data[13], data[14], data[15],
		)
	}

	return line, nil
}

func main() {
	kingpin.Version("0.1.0")
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
