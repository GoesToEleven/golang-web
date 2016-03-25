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
}

// Model returns a value of type model
func Model(c *http.Cookie, req *http.Request) model {
	xs := strings.Split(c.Value, "|")
	usrData := xs[1]

	// in cookie:
	// model encoded to JSON encode to base64
	// in memcache:
	// model encoded to JSON
	bs, err := base64.URLEncoding.DecodeString(usrData)
	if err != nil {
		log.Println("Error decoding base64", err)
	}

	m := unmarshalModel(bs)

	// if data is in memcache
	// get pictures from there
	// see refactor-notes.md for explanation
	id := xs[0]
	m2 := retrieveMemc(req, id)
	if m2.Pictures != nil {
		m.Pictures = m2.Pictures
		log.Println("Picture paths returned from memcache")
	}

	return m
}

func unmarshalModel(bs []byte) model {

	var m model
	err := json.Unmarshal(bs, &m)
	if err != nil {
		fmt.Println("error unmarshalling: ", err)
	}

	return m
}
