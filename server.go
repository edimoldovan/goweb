package main

import (
	"flag"
	"fmt"
	"log"
	"main/router"
	"net"
	"net/http"
)

func main() {
	var (
		host = flag.String("host", "", "host http address to listen on")
		port = flag.String("port", "8000", "port number for http listener")
	)
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)
	if err := runHttp(addr); err != nil {
		log.Fatal(err)
	}
}

func runHttp(listenAddr string) error {
	s := http.Server{
		Addr:    listenAddr,
		Handler: router.New(),
	}
	fmt.Printf("Starting HTTP listener at %s\n", listenAddr)
	return s.ListenAndServe()
}
