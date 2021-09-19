# healthcheck [![Build status](https://travis-ci.org/tischda/healthcheck.svg?branch=master)](https://travis-ci.org/tischda/healthcheck)

Command line utility to monitor a web service inside a docker container. The utility has 2 parameters: the
URL to monitor and a string to match in the response.

If the string is found, then the healthcheck is successful, otherwise it fails. If no string is specified,
then only the HTTP response code is checked (<400 means OK).

You use the healthcheck instruction in a dockerfile (see example below).

More information at https://docs.docker.com/engine/reference/builder/#healthcheck

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
