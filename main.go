package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var version string

var flagQuiet bool
var flagVersion bool
var flagTimeout int

func init() {
	flag.BoolVar(&flagQuiet, "quiet", false, "do not print anything to console")
	flag.BoolVar(&flagQuiet, "q", false, "do not print anything to console (shorthand)")
	flag.BoolVar(&flagVersion, "version", false, "print version and exit")
	flag.BoolVar(&flagVersion, "v", false, "print version and exit (shorthand)")
	flag.IntVar(&flagTimeout, "timeout", 30, "connection timeout in seconds")
	flag.IntVar(&flagTimeout, "t", 30, "connection timeout in seconds (shorthand)")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] <URL> [text to search for in HTTP response]\nOPTIONS:\n", os.Args[0])
		flag.PrintDefaults()
	}

	log.SetFlags(0)
	flag.Parse()

	if flag.Arg(0) == "version" || flagVersion {
		fmt.Println("healthcheck version", version)
		return
	}

	if flag.NArg() < 1 || flag.NArg() > 2 {
		flag.Usage()
		os.Exit(1)
	}
	healthCheck(flag.Arg(0), flag.Arg(1))
}

func healthCheck(url string, text string) {
	client := http.Client{
		Timeout: time.Duration(flagTimeout) * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		fatal("Connection error:", err)
	} else {
		if resp.StatusCode < 400 {
			info(url, ": Online - response:", resp.StatusCode)
			if text != "" {
				checkHttpBody(url, resp.Body, text)
			}
		} else {
			fatal(url, ": Error - response:", resp.StatusCode)
		}
	}
}

func checkHttpBody(url string, body io.ReadCloser, text string) {
	defer body.Close()
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		fatal(url, ": Error reading HTTP body:", err)
	}
	if !strings.Contains(string(bodyBytes), text) {
		fatal(url, ": Error - string not found in HTTP response:", text)
	}
}

func info(v ...interface{}) {
	if !flagQuiet {
		log.Println(v...)
	}
}

func fatal(v ...interface{}) {
	info(v...)
	os.Exit(1)
}
