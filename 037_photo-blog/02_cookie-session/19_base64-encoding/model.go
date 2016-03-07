package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"net/http"
	"encoding/base64"
	"log"
)

type model struct {
	State    bool
	Pictures []string
}

func Model(c *http.Cookie) model {
	xs := strings.Split(c.Value, "|")
	usrData := xs[1]

	// decode from base64
	bs, err := base64.URLEncoding.DecodeString(usrData)
	if err != nil {
		log.Println("Error decoding base64", err)
	}

	// unmarshal from JSON
	var m model
	err = json.Unmarshal(bs, &m)
	if err != nil {
		fmt.Println("error unmarshalling: ", err)
	}
	return m
}