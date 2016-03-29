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

	bs, err := base64.URLEncoding.DecodeString(usrData)
	if err != nil {
		log.Println("Error decoding base64", err)
	}

	m := unmarshalModel(bs)

	id := xs[0]
	m2 := retrieveMemc(req, id)
	if m2.Pictures != nil {
		m.Pictures = m2.Pictures
		log.Println("PICTURE PATHS RETURNED FROM MEMCACHE")
		log.Println(m.Pictures)
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
