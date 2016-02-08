package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main(){
	http.HandleFunc("/", index)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("public/css"))))
	http.Handle("/pic/", http.StripPrefix("/pic", http.FileServer(http.Dir("public/pic"))))
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}