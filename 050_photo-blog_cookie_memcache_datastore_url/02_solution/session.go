package mem

import (
"net/http"
	"log"
	"errors"
)

func getID(res http.ResponseWriter, req *http.Request) (string, error) {
	var id, origin string
	var cookie *http.Cookie
	// try to get the id from the COOKIE
	origin = "COOKIE"
	cookie, err := req.Cookie("session-id")
	if err == http.ErrNoCookie {
		// try to get the id from the URL
		origin = "URL"
		id := req.FormValue("id")
		if id == "" {
			// no id, so create one BRAND NEW
			origin = "BRAND NEW VIA LOGOUT"
			log.Println("ID CAME FROM", origin)
			http.Redirect(res, req, "/logout", http.StatusSeeOther)
			return id, errors.New("ERROR: redirect to /logout because no session id accessible")
		}
		// try to store id for later use in COOKIE
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	id = cookie.Value
	log.Println("ID CAME FROM", origin)
	return id, nil
}
