package main

import (
	"net/http"
	"time"
	"fmt"
	"flag"
	"os"
)

var sleepSec = flag.Int("sleep", 30, "Time (in seconds) to sleep between requests.")
var addr = flag.String("addr", "localhost:80", "ip:port or dns:port to issue requests.")

func main() {
	flag.Usage = func() {
		fmt.Printf("Repeatedly time an HTTP request.\n\n")
		fmt.Printf("Usage: %s [-sleep=SECONDS] -addr ADDRESS:PORT\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	sleepFor := time.Second * time.Duration(*sleepSec)
	fmt.Printf("Checking '%s' every %s\n", *addr, sleepFor)
	for {
		TimeIt()
		time.Sleep(sleepFor)
	}
}

func TimeIt() {
	var start, stop time.Time
	var err error
	var response *http.Response
	client := &http.Client{}

	start = time.Now()
	response, err = client.Get(*addr)
	stop = time.Now()

	durr := stop.Sub(start)
	if err != nil {
		fmt.Printf("[%s] Error during http: %s\n", durr, err)
	} else {
		fmt.Printf("[%s] %s\n", durr, response.Status)
	}
}
