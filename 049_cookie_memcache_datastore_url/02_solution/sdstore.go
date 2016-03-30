package dstore

import (
"google.golang.org/appengine"
"google.golang.org/appengine/datastore"
	"net/http"
	"log"
)

func storeDstore(m model, id string, req *http.Request) error {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Users", id, 0, nil)

	_, err := datastore.Put(ctx, key, &m)
	if err != nil {
		log.Println("ERROR storeDstore datastore.Put", err)
		return err
	}
	return nil
}

func retrieveDstore(id string, req *http.Request) (model, error) {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Users", id, 0, nil)

	var m model
	err := datastore.Get(ctx, key, &m)
	if err != nil {
		log.Println("ERROR retrieveDstore datastore.Get", err)
		return m, err
	}
	return m, nil
}
