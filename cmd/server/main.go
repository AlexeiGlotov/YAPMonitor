package main

import (
	"io"
	"net/http"
)

type gauge struct {
}

type counter struct {
}

func updateMetrics(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/update/", updateMetrics)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
