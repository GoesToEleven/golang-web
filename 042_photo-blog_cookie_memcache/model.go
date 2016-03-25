package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type model struct {
	Name     string
	State    bool
	Pictures []string
	req	*http.Request
}

func Model(s string, req *http.Request) model {
	xs := strings.Split(s, "|")
	usrData := xs[1]

	bs, err := base64.URLEncoding.DecodeString(usrData)
	if err != nil {
		log.Println("Error decoding base64", err)
	}

	var m model
	err = json.Unmarshal(bs, &m)
	if err != nil {
		fmt.Println("error unmarshalling: ", err)
	}
	m.req = req

	// if data is in memcache
	// get pictures from there
	id := xs[0]
	if retrieveMemc(req, id) != nil {

	}

	return m
}
