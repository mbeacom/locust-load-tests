package main

import (
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

func main() {
	count := uint64(0)
	tick := time.NewTicker(time.Second)
	http.ListenAndServe(":8080", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			atomic.AddUint64(&count, 1)
			select {
			case <-tick.C:
				log.Printf("Rate: %d", atomic.SwapUint64(&count, 0))
			default:
			}
		},
	))
}
