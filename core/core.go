package core

import (
	"log"
	"sync"
	"time"
)

type pending struct {
	sync.Mutex
	queue []chan string
}

func Start(listeners <-chan chan string) {
	p := new(pending)

	go func() {
		c := time.Tick(time.Duration(10) * time.Second)
		for range c {
			p.Lock()
			for _, listener := range p.queue {
				log.Println("Sending a toc")
				listener <- "toc"
			}

			if len(p.queue) > 0 {
				p.queue = make([]chan string, 0)
			}
			p.Unlock()
		}
	}()

	go func() {
		for listener := range listeners {
			p.Lock()
			p.queue = append(p.queue, listener)
			p.Unlock()
		}
	}()
}
