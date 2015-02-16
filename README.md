# gosnatch
an nzb snatcher written in go

currently in development and not yet ready for use.
this should be considered pre-alpha state!

## to build from source:

1. Install Go and setup GOPATH and GOROOT
2. Clone this repo to GOPATH/src/github.com/gosnatch/gosnatch
3. cd to GOPATH/src/github.com/gosnatch/gosnatch
4. `make run` to run
4.1 `make` to build a static binary called gosnatch.out


### OSX
build/run/compile on osx with sqlite support:

`go build -ldflags -linkmode=external` may be needed on some OSX Versions