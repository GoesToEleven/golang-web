package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

type Task struct {
	ID    int64 `datastore:"-"`
	Email string
	Text  string
}

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/todos", santos)
}

func index(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	http.ServeFile(res, req, "index.html")
}

func santos(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	tasks := make([]Task, 0)

	switch req.Method {
	case "GET":
		// get all of the tasks
		q := datastore.NewQuery("Task").Filter("Email =", u.Email)
		iterator := q.Run(ctx)
		for {
			var t Task
			key, err := iterator.Next(&t)
			if err == datastore.Done {
				break
			} else if err != nil {
				log.Errorf(ctx, "santos GET iterator.Next: %v", err)
				http.Error(res, err.Error(), 500)
				return
			}
			t.ID = key.IntID()
			tasks = append(tasks, t)
		}
		err := json.NewEncoder(res).Encode(tasks)
		if err != nil {
			log.Errorf(ctx, "santos GET json.NewEncoder: %v", err)
			return
		}
	case "POST":
		// the user is posting a new task
		var t Task
		err := json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			log.Errorf(ctx, "santos POST json.NewDecoder: %v", err)
			http.Error(res, err.Error(), 500)
			return
		}
		t.Email = u.Email
		// add to datastore
		key := datastore.NewIncompleteKey(ctx, "Task", nil)
		key, err = datastore.Put(ctx, key, &t)
		if err != nil {
			log.Errorf(ctx, "santos POST datastore.Put: %v", err)
			http.Error(res, err.Error(), 500)
			return
		}
		t.ID = key.IntID()
		// send back to user
		err = json.NewEncoder(res).Encode(t)
		if err != nil {
			log.Errorf(ctx, "santos POST json.NewEncoder: %v", err)
			return
		}
	case "DELETE":
		id, _ := strconv.ParseInt(req.FormValue("id"), 10, 64)
		if id == 0 {
			http.Error(res, "not found", 404)
			return
		}
		key := datastore.NewKey(ctx, "Task", "", id, nil)
		var t Task
		err := datastore.Get(ctx, key, &t)
		if err != nil {
			log.Errorf(ctx, "santos DELETE datastore.Get: %v", err)
			http.Error(res, err.Error(), 500)
			return
		}
		if t.Email != u.Email {
			http.Error(res, "access denied", 401)
			return
		}
		err = datastore.Delete(ctx, key)
		if err != nil {
			log.Errorf(ctx, "santos DELETE datastore.Delete: %v", err)
			return
		}
	default:
		http.Error(res, "Method Not Allowed", 405)
	}
}
