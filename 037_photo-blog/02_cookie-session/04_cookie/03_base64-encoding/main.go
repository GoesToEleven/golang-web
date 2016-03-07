package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io"
	"encoding/base64"
)

type model struct {
	State bool
	Pictures []string
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	data := foo()
	cookie, err := req.Cookie("session")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: data,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	io.WriteString(res, cookie.Value)
}

func foo() string {
	m := model{
		State: true,
		Pictures: []string{
			"one.jpg",
			"two.jpg",
			"three.jpg",
		},
	}

	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	encodeURL := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	str := base64.NewEncoding(encodeURL).EncodeToString(bs)
	return str
}