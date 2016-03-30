package dstore

import (
"net/http"
"github.com/nu7hatch/gouuid"
	"log"
)

func getID(res http.ResponseWriter, req *http.Request) (string, error) {

	var id string
	var cookie *http.Cookie
	cookie, err := req.Cookie("session-id")
	if err != nil {
		pid, err := uuid.NewV4()
		if err != nil {
			log.Println("ERROR getID uuid.NewV4", err)
			return id, err
		}
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: pid.String(),
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	id = cookie.Value
	return id, nil
}
