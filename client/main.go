package main

import (
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/relistan/nrpe-grpc/nrperpc"
)

const (
	address = "localhost:50000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := nrperpc.NewCheckClient(conn)

	if len(os.Args) < 2 {
		log.Fatal("Pass the check name as an arg")
	}

	resp, err := c.NrpeCheck(context.Background(), &nrperpc.NrpeRequest{Name: os.Args[1]})
	if err != nil {
		log.Fatalf("could not call nrpe: %v", err)
	}
	log.Printf("%v: %v", resp.StatusCode, resp.StatusLine)
}
