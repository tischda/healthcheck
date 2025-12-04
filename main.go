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

// https://goreleaser.com/cookbooks/using-main.version/
var (
	name    string
	version string
	date    string
	commit  string
)

// flags
type Config struct {
	quiet   bool
	timeout int
	help    bool
	version bool
}

func initFlags() *Config {
	cfg := &Config{}
	flag.BoolVar(&cfg.quiet, "quiet", false, "do not print anything to console")
	flag.BoolVar(&cfg.quiet, "q", false, "do not print anything to console (shorthand)")
	flag.IntVar(&cfg.timeout, "timeout", 30, "connection timeout in seconds")
	flag.IntVar(&cfg.timeout, "t", 30, "connection timeout in seconds (shorthand)")
	flag.BoolVar(&cfg.help, "?", false, "")
	flag.BoolVar(&cfg.help, "help", false, "displays this help message")
	flag.BoolVar(&cfg.version, "v", false, "")
	flag.BoolVar(&cfg.version, "version", false, "print version and exit")
	return cfg
}

func main() {
	log.SetFlags(0)
	cfg := initFlags()
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+name+` [OPTIONS] <URL> [string to look for in HTTP response]

OPTIONS:
  -q, --quiet
          do not print anything to console
  -t int
          connection timeout in seconds (shorthand) (default 30)
  -timeout int
          connection timeout in seconds (default 30)
  -?, --help
          display this help message
  -v, --version
          print version and exit

EXAMPLES:`)

		fmt.Fprintln(os.Stderr, "\n  $ "+name+` https://example.com examples
  https://example.com : Online - response: 200`)
	}
	flag.Parse()

	if flag.Arg(0) == "version" || cfg.version {
		fmt.Printf("%s %s, built on %s (commit: %s)\n", name, version, date, commit)
		return
	}

	if cfg.help {
		flag.Usage()
		return
	}

	if flag.NArg() < 1 || flag.NArg() > 2 {
		flag.Usage()
		os.Exit(1)
	}
	healthCheck(cfg, flag.Arg(0), flag.Arg(1))
}

func healthCheck(cfg *Config, url string, text string) {
	client := http.Client{
		Timeout: time.Duration(cfg.timeout) * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalln("Connection error:", err)
	} else {
		if resp.StatusCode < 400 {
			if !cfg.quiet {
				log.Println(url, ": Online - response:", resp.StatusCode)
			}
			if text != "" {
				checkHttpBody(url, resp.Body, text)
			}
		} else {
			log.Fatalln(url, ": Error - response:", resp.StatusCode)
		}
	}
}

func checkHttpBody(url string, body io.ReadCloser, text string) {
	defer body.Close()
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		log.Fatalln(url, ": Error reading HTTP body:", err)
	}
	if !strings.Contains(string(bodyBytes), text) {
		log.Fatalln(url, ": Error - string not found in HTTP response:", text)
	}
}
