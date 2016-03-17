package posts

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/memcache"
	"encoding/json"
"log"
)

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/words", showWords)
}

type Word struct {
	Term       string
	Definition string
}

func index(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path == "/favicon.ico" {
		http.NotFound(res, req)
		return
	}

	res.Header().Set("Content-Type", "text/html")

	if req.Method == "POST" {
		putWord(res, req)
	}

	fmt.Fprintln(res, `
			<form method="POST" action="/">
				<h1>Word</h1>
				<input type="text" name="term"><br>
				<h1>Definition</h1>
				<textarea name="definition"></textarea>
				<input type="submit">
			</form>`)
}

func putWord(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	term := req.FormValue("term")
	definition := req.FormValue("definition")
	entity := &Word{
		Term:       term,
		Definition: definition,
	}

	// put the word in MEMCACHE
	// item.Key is xWord
	// item.Value is []Word
	var words []Word
	item, _ := memcache.Get(ctx, "xWord")
	err := json.Unmarshal(item.Value, &words)
	if err != nil {
		log.Println(err)
	}
	words = append(words, Word{term, definition})
	bs, err := json.Marshal(words)
	if err != nil {
		log.Println(err)
	}
	item.Value = bs
	memcache.Set(ctx, &item)

	// put the word in DATASTORE
	key := datastore.NewKey(ctx, "Word", term, 0, nil)
	_, err = datastore.Put(ctx, key, entity)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
}

func showWords(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	var html string

	// get the words from MEMCACHE
	// item.Key is xWord
	// item.Value is []Word
	item, _ := memcache.Get(ctx, "xWord")
	if item.Value != "" {
		var words []Word
		err := json.Unmarshal(item.Value, &words)
		if err != nil {
			log.Println(err)
		}
		for _, v := range words {
			html += `
			<dt>` + v.Term + `</dt>
			<dd>` + v.Definition + `</dd>
		`
		}
		return
	}

	// get the words from DATASTORE
	q := datastore.NewQuery("Word").Limit(3).Order("-Term")
	ctx := appengine.NewContext(req)

	iterator := q.Run(ctx)
	for {
		var entity Word
		_, err := iterator.Next(&entity)
		if err == datastore.Done {
			break
		} else if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		html += `
			<dt>` + entity.Term + `</dt>
			<dd>` + entity.Definition + `</dd>
		`
	}

	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(res, `<dl>`+html+`</dl>`)
}
