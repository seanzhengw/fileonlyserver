# fileonlyserver

static file handler without directory listings for net/http

This is just a packaging of http.FileServer,

but disabled directory listings by return 404.

## Install

	go get -u github.com/seanzhengw/fileonlyserver

## Examples

use it like http.FileServer

	http.Handle("/", fileonlyserver.Serve(http.Dir("/tmp")))
