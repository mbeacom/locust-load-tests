package main

import (
	"fmt"
	"github.com/myzhan/boomer"
	"net/http"
)

func urlBash() {
	start := boomer.Now()

	localServer := "http://localhost:6060"
	resp, _ := http.Get(localServer)

	elapsed := boomer.Now() - start

	if resp.StatusCode == 200 {
		boomer.Events.Publish("request_success", "http", "urlBash", elapsed, int64(10))
	} else {
		boomer.Events.Publish("request_failure", "http", "urlBash", elapsed, int64(10))
	}

}

func main() {

	task := &boomer.Task{
		Name: "urlBash",
		Fn:   urlBash,
	}

	boomer.Run(task)
}
