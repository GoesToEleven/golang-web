package main

import (
	"github.com/nu7hatch/gouuid"
	"net/http"
	"io"
	"strings"
)

func main() {

	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	cookie, err := req.Cookie("session-id")

	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
	}

	if req.FormValue("email") != "" && !strings.Contains(cookie.Value, "email") {
		cookie.Value = cookie.Value + ` email=` + req.FormValue("email")
	}

	http.SetCookie(res, cookie)

	io.WriteString(res, `<!DOCTYPE html>
	<html>
	  <body>
	    <form method="POST">
	    `+cookie.Value+`
	      <input type="email" name="email">
	      <input type="submit">
	    </form>
	  </body>
	</html>`)

}