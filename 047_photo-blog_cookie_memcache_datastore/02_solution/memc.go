package mem

import (
	"encoding/base64"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"log"
	"net/http"
)

func storeMemc(bs []byte, id string, req *http.Request) {
	ctx := appengine.NewContext(req)
	item1 := memcache.Item{
		Key:   id,
		Value: bs,
	}
	memcache.Set(ctx, &item1)
}

func retrieveMemc(req *http.Request, id string) (model, error) {
	ctx := appengine.NewContext(req)
	item, err := memcache.Get(ctx, id)
	if err != nil {
		log.Println("Error retrieving memc", err)
		return nil, err
	}

	// decode item.Value from base64
	bs, err := base64.URLEncoding.DecodeString(string(item.Value))
	if err != nil {
		log.Println("Error decoding base64 in retrieveMemc", err)
		return nil, err
	}

	// unmarshal from JSON
	var m model
	if item != nil {
		m = unmarshalModel(bs)
	}
	return m, nil
}
