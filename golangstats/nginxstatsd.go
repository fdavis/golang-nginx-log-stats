package golangstats

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	KEY_STATUS_CODE  = "status_code"
	KEY_STATUS_ROUTE = "request_route"
)

type HttpStats struct {
	fiveHundreds, fourHundreds, threeHundreds, twoHundreds int
	errorUrls                                              map[string]int
}

func (stats *HttpStats) clear() {
	stats.fiveHundreds, stats.fourHundreds, stats.threeHundreds, stats.twoHundreds = 0, 0, 0, 0
	stats.errorUrls = make(map[string]int)
}

func (stats *HttpStats) update(m map[string]string) {
	status_code, err := strconv.Atoi(m[KEY_STATUS_CODE])
	check(err)
	if status_code >= 200 && status_code <= 299 {
		stats.twoHundreds++
	} else if status_code >= 300 && status_code <= 399 {
		stats.threeHundreds++
	} else if status_code >= 400 && status_code <= 499 {
		stats.fourHundreds++
	} else if status_code >= 500 && status_code <= 599 {
		stats.fiveHundreds++
		stats.errorUrls[m[KEY_STATUS_ROUTE]]++
	}
}

func (stats *HttpStats) showStats() string {
	retStr := fmt.Sprintf("50x:%d|s\n", stats.fiveHundreds) +
		fmt.Sprintf("40x:%d|s\n", stats.fourHundreds) +
		fmt.Sprintf("30x:%d|s\n", stats.threeHundreds) +
		fmt.Sprintf("20x:%d|s\n", stats.twoHundreds)
	for route, count := range stats.errorUrls {
		retStr += fmt.Sprintf("%s:%d|s\n", route, count)
	}
	return retStr
}

func parseLine(line string) map[string]string {
	ret := make(map[string]string)
	// regex throws away first part of log until the request
	// save request route and
	regex := regexp.MustCompile(`[a-zA-Z0-9 /:,.+\[\]-]+"[A-Z]+ ([a-z0-9()+,\-.:=@;$_!*'%/?#]+) HTTP/1.[01]" ([0-9]{3})`)
	matches := regex.FindStringSubmatch(line)
	ret[KEY_STATUS_CODE] = matches[1]
	ret[KEY_STATUS_ROUTE] = matches[2]
	return ret
}

func parseLogs(logFile string, stats *HttpStats, logPosition int64) int64 {
	f, err := os.Open(logFile)
	check(err)
	defer f.Close()

	f.Seek(logPosition, os.SEEK_SET)
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		s := string(line)
		result := parseLine(s)
		fmt.Println(result)
		stats.update(result)
		fmt.Println(stats.showStats())
		line, isPrefix, err = r.ReadLine()
	}

	if isPrefix {
		panic(errors.New("buffer size too small"))
	}
	if err != io.EOF {
		panic(err)
	}
	ret, err := f.Seek(logPosition, os.SEEK_SET)
	check(err)
	return ret
}

func check(e error) bool {
	if e != nil {
		panic(e)
	}
	return true
}

func main() {
	poll := flag.Bool("poll", true, "if set keep running in the foreground else parse once and quit")
	logFilename := flag.String("logFilename", "/var/log/nginx/access.log", "default access.log file to parse: /var/log/nginx/access.log")
	//	statsFilename := flag.String("statsfilename", "/var/log/stats.log", "default stats.log to write to: /var/log/stats.log")

	flag.Parse()

	var myStats HttpStats
	myStats.clear()

	fmt.Println("poll var:", *poll)
	if *poll {
		ticker := time.NewTicker(time.Second * 5)
		go func() {
			var logPosition int64 = 0
			for t := range ticker.C {
				fmt.Println("Tick at", t)
				logPosition = parseLogs(*logFilename, &myStats, logPosition)
			}
		}()
		// wait forever
		select {}
	} else {
		parseLogs(*logFilename, &myStats, 0)
		fmt.Printf(myStats.showStats())
	}
}
