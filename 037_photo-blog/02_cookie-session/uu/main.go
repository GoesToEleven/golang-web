package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":9999", nil)
}

func index(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("session")
	fmt.Printf("%T \n", cookie)
	fmt.Printf("%T \n", err)
	fmt.Println(cookie)
	fmt.Println(err)

}
