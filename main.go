package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var version string

var flagQuiet = flag.Bool("quiet", false, "do not print anything to console")
var flagVersion = flag.Bool("version", false, "print version and exit")

func init() {
	flag.BoolVar(flagQuiet, "q", false, "")
	flag.BoolVar(flagVersion, "v", false, "")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <URL> [text to search for in HTTP response]\n", os.Args[0])
		flag.PrintDefaults()
	}

	log.SetFlags(0)
	flag.Parse()

	if flag.Arg(0) == "version" || *flagVersion {
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
	resp, err := http.Get(url)
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
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		fatal(url, ": Error reading HTTP body:", err)
	}
	if !strings.Contains(string(bodyBytes), text) {
		fatal(url, ": Error - string not found in HTTP response:", text)
	}
}

func info(v ...interface{}) {
	if !*flagQuiet {
		log.Println(v...)
	}
}

func fatal(v ...interface{}) {
	if !*flagQuiet {
		log.Fatalln(v...)
	}
}
