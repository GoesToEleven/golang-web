package mem

import (
	"net/http"
	"log"
)

func getCookie(res http.ResponseWriter, req *http.Request) (*http.Cookie, error) {

	cookie, err := req.Cookie("session-id")

	if err != nil {
		cookie, err = newVisitor(req)
		if err != nil {
			log.Println("ERROR getCookie newVisitor", err)
			return nil, err
		}
		http.SetCookie(res, cookie)
		return cookie, nil
	}

	return cookie, nil
}

func makeCookie(m model, id string, req *http.Request) (*http.Cookie, error) {

	// DATASTORE
	err := storeDstore(m, id, req)
	if err != nil {
		log.Println("ERROR makeCookie storeDstore", err)
		return nil, err
	}

	// MEMCACHE
	err = storeMemc(m, id, req)
	if err != nil {
		log.Println("ERROR makeCookie storeMemc", err)
		return nil, err
	}

	// COOKIE
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie, nil
}