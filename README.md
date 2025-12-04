[![Build Status](https://github.com/tischda/healthcheck/actions/workflows/build.yml/badge.svg)](https://github.com/tischda/healthcheck/actions/workflows/build.yml)
[![Test Status](https://github.com/tischda/healthcheck/actions/workflows/test.yml/badge.svg)](https://github.com/tischda/healthcheck/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/tischda/healthcheck/badge.svg)](https://coveralls.io/r/tischda/healthcheck)
[![Linter Status](https://github.com/tischda/healthcheck/actions/workflows/linter.yml/badge.svg)](https://github.com/tischda/healthcheck/actions/workflows/linter.yml)
[![License](https://img.shields.io/github/license/tischda/healthcheck)](/LICENSE)
[![Release](https://img.shields.io/github/release/tischda/healthcheck.svg)](https://github.com/tischda/healthcheck/releases/latest)


# healthcheck

Command line utility to monitor a web service inside a docker container. The utility has 2 parameters: the
URL to monitor and a string to match in the response.

If the string is found, then the healthcheck is successful, otherwise it fails. If no string is specified,
then only the HTTP response code is checked (<400 means OK).

You could use it with the HEALTHCHECK instruction in a dockerfile:

~~~
COPY target/healthcheck /healthcheck
HEALTHCHECK CMD [ "/healthcheck",  "--quiet",  "http://localhost:8080/api", "search.bleve" ]
~~~

More information at https://docs.docker.com/engine/reference/builder/#healthcheck

## Install

~~~
go install github.com/tischda/healthcheck@latest
~~~

## Usage

~~~
Usage: healthcheck.exe [OPTIONS] <URL> [text to search for in HTTP response]
OPTIONS:
  -q    do not print anything to console (shorthand)
  -quiet
        do not print anything to console
  -t int
        connection timeout in seconds (shorthand) (default 30)
  -timeout int
        connection timeout in seconds (default 30)
  -v    print version and exit (shorthand)
  -version
        print version and exit
~~~
