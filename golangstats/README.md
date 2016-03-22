## About the go script

### go test

```bash
$ go test nginxstatsd_test.go
ok      command-line-arguments  0.008s
```
### testRun.sh

This script runs the nginxstats.go program against test\_nginx\_access.log and compares the test\_stats\_expected.log file with the generated test\_stats\_out.log file.

```bash
$ ./testRun.sh 
Files test_stats_out.log and test_stats_expected.log are identical
```

### run manually

Set `-poll=false` to run once on the command line for testing. 
In with `-debug` set every line parsed is printed out along with the key-value pairs created from the regular expression match.
You can specify what input file to grab and where to write the statsd file, see help for details.

#### help
```bash
$ go run nginxstatsd.go -h
Usage of /var/folders/rj/mj3jzcfd2z7_z_g2xmtzphkc0000gp/T/go-build721768344/command-line-arguments/_obj/exe/nginxstatsd:
  -debug
        if set log more verboseley
  -inputLogFilename string
        default access.log file to parse: /var/log/nginx/access.log (default "/var/log/nginx/access.log")
  -poll
        if set keep running in the foreground else parse once and quit (default true)
  -statsFilename string
        default stats.log to write to: /var/log/stats.log (default "/var/log/stats.log")
exit status 2
```

#### run in the terminal
```bash
$ go run nginxstatsd.go -poll=false -inputLogFilename sample.log -statsFilename stats.log
regex parse failed, skipping line:10.10.180.40 - 162.248.206.82 - - - [02/Aug/2015:17:57:37 +0000]  http http http "GET /assets/images/product-images/thumb/BGEL-3Z-CURRENT-24b30fb635781809c9ea42702e77090c.png HTTP/1.1" 200 68556 "-" "imgix/1.0" 
```

#### run in debug mode
```bash
$ go run nginxstatsd.go -poll=false -inputLogFilename sample.log -statsFilename stats.log -debug
map[status_code:200 request_route:/]
50x:0|s
40x:152|s
30x:40|s
20x:3532|s
map[request_route:/api/v1/user status_code:200]
50x:0|s
40x:152|s
30x:40|s
20x:3533|s
map[status_code:200 request_route:/api/v1/user]
50x:0|s
40x:152|s
30x:40|s
20x:3534|s
map[request_route:/api/v1/user status_code:200]
50x:0|s
40x:152|s
30x:40|s
20x:3535|s
map[status_code:200 request_route:/robots.txt]
50x:0|s
40x:152|s
30x:40|s
20x:3536|s
Could not match log line: 10.10.180.161 - 66.249.69.23, 192.33.26.238 - - - [03/Aug/2015:16:08:44 +0000]  https https https "GET /our-products/shave?action=redeem-gift-card&_escaped_fragment_= HTTP/1.1" 200 24648 "-" "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)" 
```

## Example run printing tick times, shows ~10ms jitter

```bash
$ go run nginxstatsd.go 
Tick at 2016-03-20 21:54:20.990495901 -0700 PDT
Tick at 2016-03-20 21:54:25.990719819 -0700 PDT
Tick at 2016-03-20 21:54:30.991631208 -0700 PDT
Tick at 2016-03-20 21:54:35.991436487 -0700 PDT
Tick at 2016-03-20 21:54:40.987356003 -0700 PDT
Tick at 2016-03-20 21:54:45.98700259 -0700 PDT
Tick at 2016-03-20 21:54:50.987660878 -0700 PDT
Tick at 2016-03-20 21:54:55.986779463 -0700 PDT
Tick at 2016-03-20 21:55:00.987673212 -0700 PDT
Tick at 2016-03-20 21:55:05.987777269 -0700 PDT
Tick at 2016-03-20 21:55:10.987649911 -0700 PDT
Tick at 2016-03-20 21:55:15.987691204 -0700 PDT
Tick at 2016-03-20 21:55:20.987648004 -0700 PDT
Tick at 2016-03-20 21:55:25.98769265 -0700 PDT
Tick at 2016-03-20 21:55:30.987725999 -0700 PDT
Tick at 2016-03-20 21:55:35.987423771 -0700 PDT
Tick at 2016-03-20 21:55:40.986806068 -0700 PDT
Tick at 2016-03-20 21:55:45.986843931 -0700 PDT
Tick at 2016-03-20 21:55:50.987201156 -0700 PDT
Tick at 2016-03-20 21:55:55.986955083 -0700 PDT
Tick at 2016-03-20 21:56:00.987639467 -0700 PDT
Tick at 2016-03-20 21:56:05.986942749 -0700 PDT
Tick at 2016-03-20 21:56:10.98772314 -0700 PDT
```
