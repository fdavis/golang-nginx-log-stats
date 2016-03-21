package golangstats

import "testing"
import golangstats "."

//import (
//
//"../golangnginxtstats"
//)

//var (
//	testCase1 = [...]string{"10.10.180.161 - 66.249.64.176, 192.33.28.238 - - - [02/Aug/2015:16:14:44 +0000]  http http http \"GET /robots.txt HTTP/1.1\" 404 136 \"-\" \"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)\"", "404", "/robots.txt"}
//	testCase2 = [...]string{"10.10.180.161 - 180.76.15.140, 192.33.28.238 - - - [02/Aug/2015:20:27:44 +0000]  http http http \"GET /page/privacy-policy HTTP/1.1\" 301 178 \"-\" \"Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)\"", "301", "/page/privacy-policy"}
//	testCase3 = [...]string{"10.10.180.161 - 50.112.166.232, 192.33.28.238 - - - [02/Aug/2015:15:56:14 +0000]  https https https \"GET /our-products HTTP/1.1\" 200 35967 \"-\" \"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.81 Safari/537.36\"", "200", "/our-products"}
//)

func TestParseLine(t *testing.T) {
	// Array order is log string, status code, route
	testCase1 := [...]string{"10.10.180.161 - 66.249.64.176, 192.33.28.238 - - - [02/Aug/2015:16:14:44 +0000]  http http http \"GET /robots.txt HTTP/1.1\" 404 136 \"-\" \"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)\"", "404", "/robots.txt"}
	testCase2 := [...]string{"10.10.180.161 - 180.76.15.140, 192.33.28.238 - - - [02/Aug/2015:20:27:44 +0000]  http http http \"GET /page/privacy-policy HTTP/1.1\" 301 178 \"-\" \"Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)\"", "301", "/page/privacy-policy"}
	testCase3 := [...]string{"10.10.180.161 - 50.112.166.232, 192.33.28.238 - - - [02/Aug/2015:15:56:14 +0000]  https https https \"GET /our-products HTTP/1.1\" 200 35967 \"-\" \"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.81 Safari/537.36\"", "200", "/our-products"}
	cases := []struct {
		in, want_code, want_route string
	}{
		{testCase1[0], testCase1[1], testCase1[2]},
		{testCase2[0], testCase2[1], testCase2[2]},
		{testCase3[0], testCase3[1], testCase3[2]},
	}
	for _, c := range cases {
		got := golangstats.ParseLine(c.in)
		if got[golangstats.KEY_STATUS_CODE] != c.want_code || got[golangstats.KEY_STATUS_ROUTE] != c.want_route {
			t.Errorf("ParseLine(%s)\nyields %v\nwant status %d, route %s", c.in, got, c.want_code, c.want_route)
		}
	}
}
