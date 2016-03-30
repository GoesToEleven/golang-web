package mem

import (
"net/http"
"github.com/nu7hatch/gouuid"
	"log"
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
			origin = "BRAND NEW"
			pid, err := uuid.NewV4()
			if err != nil {
				log.Println("ERROR getID uuid.NewV4", err)
				return id, err
			}
			id = pid.String()
		}
		// try storing id in a cookie for later use
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
