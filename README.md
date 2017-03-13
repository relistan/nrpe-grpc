NRPE GRPC PoC
=============

This is a demo of how you might bridge between a high-throughput Go server
and NRPE checks without compromising the Go server by compiling it against
OpenSSL which is required for NRPE's old ADH cipher suites.

Build the server with `go build` and the client with `cd client && go build`

The remote NRPE server name is hardcoded currently.
