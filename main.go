package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/finn-ball/optic-yellow-server/pkg/server"
)

var domain string
var port uint
var headless bool

func main() {
	// parse flags
	cmd()
	// create listen
	addr := fmt.Sprintf("%s:%d", domain, port)
	fmt.Println("server starting...")
	fmt.Println("domain:", addr)
	fmt.Println("headless: ", headless)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	// start server
	s := server.NewServer(headless)
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}

// cmd will parse the user input.
func cmd() {
	flag.StringVar(&domain, "d", "localhost", `Domain`)
	flag.UintVar(&port, "p", 8080, `Port to listen`)
	flag.BoolVar(&headless, "headless", false, `Run in headless mode`)
	flag.Parse()
}
