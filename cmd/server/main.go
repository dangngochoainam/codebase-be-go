package main

import (
	"example/pkg/httpserver"
	"example/pkg/ws"
	"flag"
	"log"
)

func init() {
	log.Println("Init application...")
}

func main() {

	server := flag.String("server", "", "websocket")
	flag.Parse()

	if *server == "websocket" {
		ws.StartWebsocketServer()
		return
	}
	httpserver.StartHTTPServer()
}
