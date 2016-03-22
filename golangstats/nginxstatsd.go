package main

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

var debug bool = false

func debugOut(s string) {
	if debug {
		fmt.Println(s)
	}
}

func check(e error) bool {
	if e != nil {
		panic(e)
	}
	return true
}

// Http stats counter type with helper functions
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

// regex extract stats code and request route
func ParseLine(line string) (map[string]string, error) {
	ret := make(map[string]string)
	// regex throws away first part of log until the request and ignores everything past status code
	// save request route and status code
	regex := regexp.MustCompile(`[a-zA-Z0-9 /:,.+\[\]-]+"[A-Z]+ ([a-z0-9()+,\-.:=@;$_!*'%/?#]+) HTTP/1.[01]" ([0-9]{3})`)
	matches := regex.FindStringSubmatch(line)
	if matches == nil {
		debugOut(fmt.Sprintf("Could not match log line: %s\n", line))
		return ret, errors.New("skipping line: regex parse log line failed")
	}
	ret[KEY_STATUS_ROUTE] = matches[1]
	ret[KEY_STATUS_CODE] = matches[2]
	return ret, nil
}

// read to end of nginx log file and return the offset
func parseLogs(logFile string, stats *HttpStats, logPosition int64) int64 {
	f, err := os.Open(logFile)
	check(err)
	defer f.Close()

	f.Seek(logPosition, os.SEEK_SET)
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	// todo test moving update into for loop
	for err == nil && !isPrefix {
		s := string(line)
		var result map[string]string
		var parseErr error
		result, parseErr = ParseLine(s)
		if parseErr == nil {
			debugOut(fmt.Sprintln(result))
			stats.update(result)
			debugOut(fmt.Sprintln(stats.showStats()))
		}
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

func main() {
	poll := flag.Bool("poll", true, "if set keep running in the foreground else parse once and quit")
	flag.BoolVar(&debug, "debug", false, "if set log more verboseley")
	inputLogFilename := flag.String("inputLogFilename", "/var/log/nginx/access.log", "default access.log file to parse: /var/log/nginx/access.log")
	// outputLogFilename := flag.String("outputLogFilename", "nginxstats.log", "default log destination file: nginxstats.log")
	statsFilename := flag.String("statsFilename", "/var/log/stats.log", "default stats.log to write to: /var/log/stats.log")

	flag.Parse()

	var myStats HttpStats
	myStats.clear()

	statsFile, err := os.OpenFile(*statsFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	check(err)
	defer statsFile.Close()

	if *poll {
		ticker := time.NewTicker(time.Second * 5)
		go func() {
			var logPosition int64 = 0
			// todo, should open source log file here, otherwise if log rotates we will seek past the end of file
			for t := range ticker.C {
				debugOut(fmt.Sprintln("Tick at", t))
				logPosition = parseLogs(*inputLogFilename, &myStats, logPosition)
				statsFile.WriteString(myStats.showStats())
				myStats.clear()
			}
		}()
		// wait forever
		select {}
	} else {
		parseLogs(*inputLogFilename, &myStats, 0)
		debugOut(fmt.Sprintf(myStats.showStats()))
		statsFile.WriteString(myStats.showStats())
	}
}
