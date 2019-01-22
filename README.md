# awslb-log-viewer

## Installation

Download from https://github.com/tkuchiki/awslb-log-viewer/releases

## Usage

```console
$ ./awslb-log-viewer --help
usage: awslb-log-viewer [<flags>]

Flags:
      --help           Show context-sensitive help (also try --help-long and --help-man).
  -f, --file=FILE      logfile
  -t, --lb-type="ALB"  ALB or ELB
      --version        Show application version.
```

## Examples

### ALB

```console
$ cat examples/alb.log | ./awslb-log-viewer | jq .
{
  "type": "h2",
  "timestamp": "2017-07-24T09:02:54.663814Z",
  "elb": "app/test-alb/xxxxxxxxxx",
  "client": "192.0.2.10:50072",
  "target": "192.0.2.100:80",
  "request_processing_time": "-1",
  "target_processing_time": "-1",
  "response_processing_time": "-1",
  "elb_status_code": "502",
  "target_status_code": "-",
  "received_bytes": "248",
  "sent_bytes": "610",
  "method": "GET",
  "http_version": "HTTP/2.0",
  "protocol": "https",
  "host": "example.com",
  "port": "443",
  "uri": "/",
  "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
  "ssl_cipher": "ECDHE-RSA-AES128-GCM-SHA256",
  "ssl_protocol": "TLSv1.2",
  "target_group_arn": "arn:aws:elasticloadbalancing:us-east-1:123456789012:targetgroup/testgroup/xxxxxxxxxx",
  "trace_id": "Root=1-xxxxxxxx-xxxxxxxxxxxx",
  "domain_name": "example.com",
  "chosen_cert_arn": "arn:aws:acm:us-east-1:123456789012:certificate/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  "matched_rule_priority": "2",
  "request_creation_time": "2019-01-21T07:32:28.468000Z",
  "actions_executed": "forward",
  "redirect_url": "-",
  "error_reason": "-"
}
{
  "type": "h2",
  "timestamp": "2017-07-24T09:03:54.663814Z",
  "elb": "app/test-alb/xxxxxxxxxx",
  "client": "192.0.2.10:50072",
  "target": "192.0.2.100:80",
  "request_processing_time": "-1",
  "target_processing_time": "-1",
  "response_processing_time": "-1",
  "elb_status_code": "502",
  "target_status_code": "-",
  "received_bytes": "248",
  "sent_bytes": "610",
  "method": "GET",
  "http_version": "HTTP/2.0",
  "protocol": "https",
  "host": "example.com",
  "port": "443",
  "uri": "/",
  "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
  "ssl_cipher": "ECDHE-RSA-AES128-GCM-SHA256",
  "ssl_protocol": "TLSv1.2",
  "target_group_arn": "arn:aws:elasticloadbalancing:us-east-1:123456789012:targetgroup/testgroup/xxxxxxxxxx",
  "trace_id": "Root=1-xxxxxxxx-xxxxxxxxxxxx",
  "domain_name": "example.com",
  "chosen_cert_arn": "arn:aws:acm:us-east-1:123456789012:certificate/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  "matched_rule_priority": "2",
  "request_creation_time": "2019-01-21T07:32:28.468000Z",
  "actions_executed": "forward",
  "redirect_url": "-",
  "error_reason": "-"
}
```

### ELB

```console
$ cat examples/elb.log | ./awslb-log-viewer -t ELB | jq .
{
  "timestamp": "2017-07-24T09:02:54.663814Z",
  "elb": "app/test-alb/xxxxxxxxxx",
  "client": "192.0.2.10:50072",
  "target": "192.0.2.100:80",
  "request_processing_time": "-1",
  "backend_processing_time": "-1",
  "response_processing_time": "-1",
  "elb_status_code": "502",
  "backend_status_code": "-",
  "received_bytes": "248",
  "sent_bytes": "610",
  "reqeust": "GET https://example.com:443/ HTTP/2.0",
  "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
  "ssl_cipher": "ECDHE-RSA-AES128-GCM-SHA256",
  "ssl_protocol": "TLSv1.2"
}
{
  "timestamp": "2017-07-24T09:03:54.663814Z",
  "elb": "app/test-alb/xxxxxxxxxx",
  "client": "192.0.2.10:50072",
  "target": "192.0.2.100:80",
  "request_processing_time": "-1",
  "backend_processing_time": "-1",
  "response_processing_time": "-1",
  "elb_status_code": "502",
  "backend_status_code": "-",
  "received_bytes": "248",
  "sent_bytes": "610",
  "reqeust": "GET https://example.com:443/ HTTP/2.0",
  "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
  "ssl_cipher": "ECDHE-RSA-AES128-GCM-SHA256",
  "ssl_protocol": "TLSv1.2"
}
```
