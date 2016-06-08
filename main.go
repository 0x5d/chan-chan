package main

import (
	"log"

	"github.com/castillobg/chan-chan/api"
	"github.com/castillobg/chan-chan/core"
)

func main() {
	listeners := make(chan chan string)
	port := 8080

	core.Start(listeners)
	log.Printf("Tic, toc, tic, toc. Listening on %d", port)
	api.Listen(8080, listeners)
}
