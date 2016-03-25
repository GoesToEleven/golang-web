package mem

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
"google.golang.org/appengine"
"google.golang.org/appengine/memcache"
)

func newVisitor() *http.Cookie {
	m, mm := initialModel()
	id, _ := uuid.NewV4()
	return makeCookie(m, mm, id.String())
}

func currentVisitor(m model, id string) *http.Cookie {
	mm := marshalModel(m)
	return makeCookie(m, mm, id)
}

func makeCookie(m model, mm []byte, id string) *http.Cookie {
	b64 := base64.URLEncoding.EncodeToString(mm)
	code := getCode(b64)
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id + "|" + b64 + "|" + code,
		// Secure: true,
		HttpOnly: true,
	}

	// send data to be stored in memcache
	storeMemc(m, mm, id)

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

func storeMemc(m model, bs []byte, id string) {
	ctx := appengine.NewContext(m.req)

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
		m = unmarshalModel(item.Value)
	}
	return m
}

func initialModel() (model, []byte) {
	m := model{
		Name:  "",
		State: false,
		Pictures: []string{
			"one.jpg",
		},
	}
	return m, marshalModel(m)
}
