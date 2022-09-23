package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Retrieving current directory.
	execFile, execError := os.Executable()
	if execError != nil {
		log.Panic(execError)
	}
	progectPath := filepath.Dir(execFile)

	// Handler functions.
	mainPage := func(w http.ResponseWriter, _ *http.Request) {
		fileData, fileError := ioutil.ReadFile(progectPath + "/index.html")
		if fileError != nil {
			log.Panic(fileError)
		}
		w.Header().Add("Content-Type", "text/html")
		w.Write(fileData)
	}

	http.HandleFunc("/", mainPage)
	fmt.Println("Server is running.")
	serverError := http.ListenAndServe(":8989", nil)

	if serverError != nil {
		log.Fatal(serverError)
	}
}
