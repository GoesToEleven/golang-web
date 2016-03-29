package mem

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"net/http"
	"log"
)

func newVisitor(req *http.Request) *http.Cookie {
	mm := initialModel()
	id, _ := uuid.NewV4()
	return makeCookie(mm, id.String(), req)
}

func currentVisitor(m model, id string, req *http.Request) *http.Cookie {
	mm := marshalModel(m)
	return makeCookie(mm, id, req)
}

func makeCookie(mm []byte, id string, req *http.Request) *http.Cookie {

	// Anytime a cookie is created, let's print the id
	// The id is the key for the value in memcache
	// Having the id will allow us to lookup the value in memcache
	log.Println("ID:", id)

	b64 := base64.URLEncoding.EncodeToString(mm)

	// SEND DATA TO BE STORED IN MEMCACHE
	storeMemc([]byte(b64), id, req)

	// SEND DATA TO BE STORED IN COOKIE
	code := getCode(b64) // hmac
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id + "|" + b64 + "|" + code,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie
}

func marshalModel(m model) []byte {
	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return bs
}

func storeMemc(bs []byte, id string, req *http.Request) {
	ctx := appengine.NewContext(req)
	item1 := memcache.Item{
		Key:   id,
		Value: bs,
	}
	memcache.Set(ctx, &item1)
}

func retrieveMemc(req *http.Request, id string) model {
	ctx := appengine.NewContext(req)
	item, _ := memcache.Get(ctx, id)

	// decode item.Value from base64
	bs, err := base64.URLEncoding.DecodeString(string(item.Value))
	if err != nil {
		log.Println("Error decoding base64 in retrieveMemc", err)
	}

	// unmarshal from JSON
	var m model
	if item != nil {
		m = unmarshalModel(bs)
	}
	return m
}

func initialModel() []byte {
	m := model{
		Name:  "",
		State: false,
		Pictures: []string{
			"one.jpg",
		},
	}
	return marshalModel(m)
}
