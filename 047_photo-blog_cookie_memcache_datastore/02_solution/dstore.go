package mem

import (
"google.golang.org/appengine"
"google.golang.org/appengine/datastore"
	"net/http"
	"log"
)

func storeDstore(m model, id string, req *http.Request) {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Photos", id, 0, nil)

	_, err := datastore.Put(ctx, key, &m)
	if err != nil {
		log.Println("ERROR PUTTING IN DATASTORE storeDstore", err)
		return
	}
	return
}

func retrieveDstore(id string, req *http.Request) (model, error) {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Photos", id, 0, nil)

	var m model
	err := datastore.Get(ctx, key, &m)
	if err == datastore.ErrNoSuchEntity {
		log.Println("NO DSTORE RETRIEVE", err)
		return nil, err
	} else if err != nil {
		log.Println("ERR DSTORE RETRIEVE", err)
		return nil, err
	}
	return m, nil
}
