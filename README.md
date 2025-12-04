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
Usage: healthcheck [OPTIONS] <URL> [string to look for in HTTP response]

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
~~~


## Examples

~~~
$ healthcheck https://example.com examples
https://example.com : Online - response: 200
~~~

~~~
$ healthcheck https://example.com notthere
https://example.com : Online - response: 200
https://example.com : Error - string not found in HTTP response: notthere
~~~

~~~
healthcheck  http://test.com
Connection error: Get "http://test.com": EOF
~~~

