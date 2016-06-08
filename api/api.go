package api

import (
	"fmt"
	"log"
	"net/http"
)

func Listen(port int, listeners chan<- chan string) {
	http.HandleFunc("/tic", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received a tic")

		message := make(chan string)
		listeners <- message
		w.Write([]byte(<-message))
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
