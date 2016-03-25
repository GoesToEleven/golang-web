package mem

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"net/http"
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
	b64 := base64.URLEncoding.EncodeToString(mm)
	code := getCode(b64)
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id + "|" + b64 + "|" + code,
		// Secure: true,
		HttpOnly: true,
	}

	// send data to be stored in memcache
	storeMemc(mm, id, req)

	// send data to be stored in a cookie
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
	// production code should not ignore the error
}

func retrieveMemc(req *http.Request, id string) model {
	ctx := appengine.NewContext(req)
	item, _ := memcache.Get(ctx, id)
	var m model
	if item != nil {
		m = unmarshalModel(string(item.Value))
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
