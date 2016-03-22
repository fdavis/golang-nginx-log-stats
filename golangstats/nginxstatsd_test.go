package main

import "testing"
import golangstats "."

func TestParseLine(t *testing.T) {
	// Array order is log string, status code, route
	// TODO, DRY this loading pattern
	testCase1 := [...]string{"10.10.180.161 - 66.249.64.176, 192.33.28.238 - - - [02/Aug/2015:16:14:44 +0000]  http http http \"GET /robots.txt HTTP/1.1\" 404 136 \"-\" \"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)\"", "404", "/robots.txt"}
	testCase2 := [...]string{"10.10.180.161 - 180.76.15.140, 192.33.28.238 - - - [02/Aug/2015:20:27:44 +0000]  http http http \"GET /page/privacy-policy HTTP/1.1\" 301 178 \"-\" \"Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)\"", "301", "/page/privacy-policy"}
	testCase3 := [...]string{"10.10.180.161 - 50.112.166.232, 192.33.28.238 - - - [02/Aug/2015:15:56:14 +0000]  https https https \"GET /our-products HTTP/1.1\" 200 35967 \"-\" \"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.81 Safari/537.36\"", "200", "/our-products"}
	// TODO: should make this case verify an error is returned
	testCase4 := [...]string{"failure to parsify", "", ""}
	testCase5 := [...]string{"66.249.67.114 - 66.249.67.114, 192.33.28.238, 66.249.67.114,127.0.0.1 - - - [03/Aug/2015:04:07:30 +0000]  http http,http http,http \"GET /shaver_images/visit.gif?x=1438246100&y=6f91013824156c47e6c93957b32a708c2835b02b HTTP/1.1\" 200 54 \"-\" \"Googlebot-Image/1.0\"", "200", "/shaver_images/visit.gif?x=1438246100&y=6f91013824156c47e6c93957b32a708c2835b02b"}
	testCase6 := [...]string{"10.10.180.40 - 162.248.206.82 - - - [02/Aug/2015:17:57:37 +0000]  http http http \"GET /assets/images/product-images/thumb/BGEL-3Z-CURRENT-24b30fb635781809c9ea42702e77090c.png HTTP/1.1\" 200 68556 \"-\" \"imgix/1.0\"", "200", "/assets/images/product-images/thumb/BGEL-3Z-CURRENT-24b30fb635781809c9ea42702e77090c.png"}
	cases := []struct {
		in, want_code, want_route string
	}{
		{testCase1[0], testCase1[1], testCase1[2]},
		{testCase2[0], testCase2[1], testCase2[2]},
		{testCase3[0], testCase3[1], testCase3[2]},
		{testCase4[0], testCase4[1], testCase4[2]},
		{testCase5[0], testCase5[1], testCase5[2]},
		{testCase6[0], testCase6[1], testCase6[2]},
	}
	for _, c := range cases {
		got, _ := golangstats.ParseLine(c.in)
		if got[golangstats.KEY_STATUS_CODE] != c.want_code || got[golangstats.KEY_STATUS_ROUTE] != c.want_route {
			t.Errorf("ParseLine(%s)\nyields %v\nwant status %d, route %s", c.in, got, c.want_code, c.want_route)
		}
	}
}
