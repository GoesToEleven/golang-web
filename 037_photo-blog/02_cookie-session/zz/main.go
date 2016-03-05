package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
"github.com/nu7hatch/gouuid"
"strings"
)

var tpl *template.Template

func init() {
	tpl, _ = template.ParseGlob("templates/*.html")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/assets/imgs/", http.StripPrefix("/assets/imgs", http.FileServer(http.Dir("./assets/imgs"))))
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

	// authenticate
	if session.Values["loggedin"] == "false" || session.Values["loggedin"] == nil {
		http.Redirect(res, req, "/login", 302)
		return
	}
	// upload photo
	src, hdr, err := req.FormFile("data")
	if req.Method == "POST" && err == nil {
		uploadPhoto(src, hdr, session)
	}
	// save session
	session.Save(req, res)
	// get photos
	data := getPhotos(session)
	// execute template
	tpl.ExecuteTemplate(res, "index.html", data)
}

func logout(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	session.Values["loggedin"] = "false"
	session.Save(req, res)
	http.Redirect(res, req, "/login", 302)
}

func login(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	if req.Method == "POST" && req.FormValue("password") == "secret" {
		session.Values["loggedin"] = "true"
		session.Save(req, res)
		http.Redirect(res, req, "/", 302)
		return
	}
	// execute template
	tpl.ExecuteTemplate(res, "login.html", nil)
}

func uploadPhoto(src multipart.File, hdr *multipart.FileHeader, session *sessions.Session) {
	defer src.Close()
	fName := getSha(src) + ".jpg"
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "assets", "imgs", fName)
	dst, _ := os.Create(path)
	defer dst.Close()
	src.Seek(0, 0)
	io.Copy(dst, src)
	addPhoto(fName, session)
}

func getSha(src multipart.File) string {
	h := sha1.New()
	io.Copy(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func addPhoto(fName string, session *sessions.Session) {
	data := getPhotos(session)
	data = append(data, fName)
	bs, _ := json.Marshal(data)
	session.Values["data"] = string(bs)
}

func getPhotos(session *sessions.Session) []string {
	var data []string
	jsonData := session.Values["data"]
	if jsonData != nil {
		json.Unmarshal([]byte(jsonData.(string)), &data)
	}
	return data
}
