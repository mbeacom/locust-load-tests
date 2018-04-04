package main

import (
	"github.com/myzhan/boomer"
	"net/http"
)

func urlBash() {
	start := boomer.Now()

	localServer := "http://localhost:6060"
	resp, _ := http.Get(localServer)

	elapsed := boomer.Now() - start

	boomer.Events.Publish(resp.Status, "http", "urlBash", elapsed, int64(10))
}

func main() {

	task := &boomer.Task{
		Name: "urlBash",
		Fn:   urlBash,
	}

	boomer.Run(task)
}
