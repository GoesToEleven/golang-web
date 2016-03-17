package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
"log"
)

func main() {
	resp, err := http.Get("http://www.golang.org")
	if err != nil {
		log.Println("http get error:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("readall error:", err)
	}

	fmt.Println(string(body))
}