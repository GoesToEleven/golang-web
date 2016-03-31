package dstore

import (
	"fmt"
	"net/http"
	"log"
)

type model struct {
	Fname string
}

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/retrieve", noConfusion)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func index(res http.ResponseWriter, req *http.Request) {

	id, err := getID(res, req)
	if err != nil {
		log.Println("ERROR index getID", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// DATASTORE
	m := model{
		Fname:"Todd",
	}
	err = storeDstore(m, id, req)
	if err != nil {
		log.Println("ERROR index storeDstore", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// MEMCACHE
	err = storeMemc(m, id, req)
	if err != nil {
		log.Println("ERROR index storeMemc", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, `
	<h1>EVERYTHING SET ID: ` + id + `</h1>
	<h1><a href="/?id=` + id + `">HOME AGAIN</a></h1>
	<h1><a href="/retrieve?id=` + id + `">RETRIEVE</a></h1>
	`)
}

func noConfusion(res http.ResponseWriter, req *http.Request) {

	id, err := getID(res, req)
	if err != nil {
		log.Println("ERROR index getID", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	html := `
	<form method="POST">
	    HIDDEN FIELD
	    <input type="hidden" name="panda" value="`+id+`">
	    <input type="submit" value="submit">
	</form>
	`

	if req.Method == "POST" {
		id = req.FormValue("panda")

		// get value
		m, err := retrieveMemc(id, req)
		if err != nil {
			log.Println("ERROR noConfusion retrieveMemc", err)
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		html += `
			<br>
			<h1>
				<a href="/?id=` + id + `">
				Go home again, `+ m.Fname +`
				</a>
			</h1>
		`
	}
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(res, html)
}
