package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	response, responseError := http.Get("http://127.0.0.1:8989/")
	if responseError != nil {
		log.Panic(responseError)
	}

	data, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		log.Panic(readError)
	}

	fmt.Println(string(data))
}
