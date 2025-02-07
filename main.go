package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":5512", nil)
	if err != nil {
		panic(err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is my service!")
}
