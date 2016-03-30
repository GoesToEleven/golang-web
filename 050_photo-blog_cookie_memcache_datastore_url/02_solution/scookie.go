package mem

import (
	"net/http"
	"log"
)

func makeCookie(m model, req *http.Request) (*http.Cookie, error) {

	// DATASTORE
	err := storeDstore(m, req)
	if err != nil {
		log.Println("ERROR makeCookie storeDstore", err)
		return nil, err
	}

	// MEMCACHE
	err = storeMemc(m, req)
	if err != nil {
		log.Println("ERROR makeCookie storeMemc", err)
		return nil, err
	}

	// COOKIE
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: m.ID,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie, nil
}