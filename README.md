# Remote shell [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/taylorflatt/remote-shell) [![Build Status](https://travis-ci.org/taylorflatt/remote-shell.svg?branch=master)](https://travis-ci.org/taylorflatt/remote-shell)
A client-server remote shell implemented in Go using gRPC and Protocol Buffers.

## Usage
* Get the source code by running "go get github.com/taylorflatt/remote-shell"
* First, run the server by going to the Server directory and typing `go run server.go`.
* Next, run the client by going to the Client directory and typing `go run client.go`. 
* In the client, enter in the server IP to establish a connection over gRPC.
* Input any shell command to be run on the server.

## Notes
To disconnect from the server, press ctrl+c or type exit (hit enter) and the client will disconnect from the server.

This client/server assumes a 12021 port.
