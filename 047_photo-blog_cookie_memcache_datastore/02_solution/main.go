package mem

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl, _ = template.ParseGlob("templates/*.html")

	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/imgs/", fs)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {

	cookie, err := getCookie(res, req)
	if err != nil {
		// problem retrieving cookie
		log.Println("ERROR index getCookie", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	id := cookie.Value

	if req.Method == "POST" {
		src, _, err := req.FormFile("data")
		if err != nil {
			log.Println("ERROR index req.FormFile", err)
			// TODO: create error page to show user
			http.Redirect(res, req, "/", http.StatusSeeOther)
			return
		}
		err = uploadPhoto(src, id, req)
		if err != nil {
			log.Println("ERROR index uploadPhoto", err)
			// expired cookie may exist on client
			http.Redirect(res, req, "/logout", http.StatusSeeOther)
			return
		}
	}

	m, err := retrieveMemc(id, req)
	if err != nil {
		log.Println("ERROR index retrieveMemc", err)
		// expired cookie may exist on client
		http.Redirect(res, req, "/logout", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "index.html", m)
}

func logout(res http.ResponseWriter, req *http.Request) {
	cookie, err := newVisitor(req)
	if err != nil {
		log.Println("ERROR logout getCookie", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(res, cookie)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func login(res http.ResponseWriter, req *http.Request) {

	cookie, err := getCookie(res, req)
	if err != nil {
		log.Println("ERROR login getCookie", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	id := cookie.Value

	if req.Method == "POST" && req.FormValue("password") == "secret" {
		m, err := retrieveMemc(id, req)
		if err != nil {
			log.Println("ERROR index retrieveMemc", err)
			// expired cookie may exist on client
			http.Redirect(res, req, "/logout", http.StatusSeeOther)
			return
		}
		m.State = true
		m.Name = req.FormValue("name")

		cookie, err := currentVisitor(m, id, req)
		if err != nil {
			log.Println("ERROR login currentVisitor", err)
			http.Redirect(res, req, "/", http.StatusSeeOther)
			return
		}
		http.SetCookie(res, cookie)

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "login.html", nil)
}
