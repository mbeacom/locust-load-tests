package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/myzhan/boomer"
)

func urlBash() {
	start := boomer.Now()

	localServer := "http://localhost:8080"
	response, err := http.Get(localServer)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	defer response.Body.Close()

	elapsed := boomer.Now() - start

	if response.StatusCode == 200 {
		boomer.Events.Publish("request_success", "http", "urlBash", elapsed, response.ContentLength)
	} else {
		boomer.Events.Publish("request_failure", "http", "urlBash", elapsed, strconv.Itoa(response.StatusCode))
	}

}

func main() {

	task := &boomer.Task{
		Name: "urlBash",
		Fn:   urlBash,
	}
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 2000

	boomer.Run(task)
}
