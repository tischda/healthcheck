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

const PROG_NAME string = "healthcheck"

// http://technosophos.com/2014/06/11/compile-time-string-in-go.html
var version string

// command line flags
var showVersion bool
var httpURL string
var str string

func init() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
	flag.StringVar(&httpURL, "URL", "", "URL of the server to check")
	flag.StringVar(&str, "string", "", "string to search for in the HTTP response")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if showVersion {
		fmt.Printf("%s version %s\n", PROG_NAME, version)
	} else {
		if httpURL == "" {
			flag.Usage()
			os.Exit(1)
		}
		healthCheck(httpURL, str)
	}
}

func healthCheck(url string, text string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("Connection error:", err)
	} else {
		if resp.StatusCode < 400 {
			log.Println(url, ": Online - response:", resp.StatusCode)
			if text != "" {
				checkHttpBody(resp.Body, text)
			}
		} else {
			log.Fatalln(url, ": Error - response:", resp.StatusCode)
		}
	}
}

func checkHttpBody(body io.ReadCloser, s string) {
	defer body.Close()
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatalln(httpURL, "Error reading HTTP body:", err)
	}
	if !strings.Contains(string(bodyBytes), s) {
		log.Fatalln(httpURL, "Error - string not found in HTTP response:", s)
	}
}
