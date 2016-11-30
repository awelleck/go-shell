# go-shell
[![Build Status](https://travis-ci.org/awelleck/go-shell.svg?branch=master)](https://travis-ci.org/awelleck/go-shell)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)

Simple shell in Go. It is assumed that this will only run on Unix based operating systems.

## Steps to run
* You will need to have [Go](https://golang.org/) installed
* Clone code: `git clone https://github.com/awelleck/go-shell.git`
* Run: `go run shell.go`

## Steps to run tests
* `cd awelleck/go-shell`
* `go test ./...`

## Lessons learned
* Go routines: call `helper` function using go routine, will execute concurrently with the calling function.
* Go channels: had to use a channel to handle SIGINT from user.
* Go concurrency patterns

## TODO
1. Find alternative to exec bash
2. Create test case for exit
