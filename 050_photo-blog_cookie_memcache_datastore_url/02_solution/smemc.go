package mem

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"log"
	"net/http"
	"encoding/json"
)

func storeMemc(m model, req *http.Request) error {
	ctx := appengine.NewContext(req)
	bs, err := json.Marshal(m)
	if err != nil {
		log.Println("ERROR storeMemc json.Marshal: ", err)
		return err
	}
	item1 := memcache.Item{
		Key:   m.ID,
		Value: bs,
	}
	err = memcache.Set(ctx, &item1)
	if err != nil {
		log.Println("ERROR storeMemc memcache.Set: ", err)
		return err
	}

	return nil
}

func retrieveMemc(id string, req *http.Request) (model, error) {

	var m model
	ctx := appengine.NewContext(req)
	item, err := memcache.Get(ctx, id)
	if err != nil {
		// get data from datastore
		m, err = retrieveDstore(id, req)
		if err != nil {
			return m, err
		}
		// put data in memcache
		storeMemc(m, req)
		return m, nil
	}
	// unmarshal from JSON
	err = json.Unmarshal(item.Value, &m)
	if err != nil {
		log.Println("ERROR retrieveMemc unmarshal", err)
		return m, err
	}
	return m, nil
}
