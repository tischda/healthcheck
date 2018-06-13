# healthcheck [![Build status](https://travis-ci.org/tischda/healthcheck.svg?branch=master)](https://travis-ci.org/tischda/healthcheck)

Utility written in [Go](https://www.golang.org).

### Install

There are no dependencies.

~~~
go get github.com/tischda/healthcheck
~~~

### Usage

~~~
Usage: ./healthcheck <URL> [text to search for in HTTP response]
  -q
  -quiet
    	do not print anything to console
  -v
  -version
    	print version and exit
~~~

Example:

~~~
COPY target/healthcheck /healthcheck
HEALTHCHECK CMD [ "/healthcheck",  "--quiet",  "http://localhost:8080/api", "search.bleve" ]
~~~
