package main

import (
	"html/template"
	"net/http"
"github.com/nu7hatch/gouuid"
"strings"
)

var tpl *template.Template

func init() {
	tpl, _ = template.ParseGlob("templates/*.html")
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")

	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String() + "|",
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}

	xs := strings.Split(cookie.Value, "|")
	token := xs[0]
	data := xs[1]


	tpl.ExecuteTemplate(res, "index.html", data)
}