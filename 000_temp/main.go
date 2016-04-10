package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, `
	<h1>Wasssup</h1>
	<p>The sky</p>
	`)
}
