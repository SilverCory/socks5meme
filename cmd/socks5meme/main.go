package main

import (
	"flag"
	"github.com/armon/go-socks5"
)

func main() {

	bindAddrPtr := flag.String("bindAddr", ":8080", "The address and port to bind to")
	flag.Parse()

	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", *bindAddrPtr); err != nil {
		panic(err)
	}

}