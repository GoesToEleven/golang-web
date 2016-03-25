package mem

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

func Model(c *http.Cookie, req *http.Request) model {
	xs := strings.Split(c.Value, "|")
	usrData := xs[1]

	m := unmarshalModel(usrData)
	m.req = req

	// if data is in memcache
	// get pictures from there
	// see refactor-notes.md for explanation
	id := xs[0]
	m2 := retrieveMemc(req, id)
	if m2.Pictures != "" {
		m.Pictures = m2.Pictures
		log.Println("Picture paths returned from memcache")
	}

	return m
}

func unmarshalModel(s string) model {

	bs, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		log.Println("Error decoding base64", err)
	}

	var m model
	err = json.Unmarshal(bs, &m)
	if err != nil {
		fmt.Println("error unmarshalling: ", err)
	}

	return m
}