package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	poll := flag.Bool("poll", false, "if set keep running in the foreground else parse once and quit")

	flag.Parse()

	fmt.Println("poll var:", *poll)
	if *poll {
		ticker := time.NewTicker(time.Second * 5)
		go func() {
			for t := range ticker.C {
				fmt.Println("Tick at", t)
			}
		}()
		// wait forever
		select {}
	} else {
		fmt.Println("No poll, parse once")
	}
}
