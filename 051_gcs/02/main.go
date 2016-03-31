package skyhdd

import (
	"net/http"
)

const gcsBucket = "learning-1130.appspot.com"

func init() {
	http.HandleFunc("/", handler)
}

func handler(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	html := `
	    <form method="POST" enctype="multipart/form-data">
		<input type="file" name="dahui">
		<input type="submit">
	    </form>
	`

	if req.Method == "POST" {
		src, hdr, err := req.FormFile("dahui")
		if err != nil {
			log.Println("error uploading photo: ", err)
			// TODO: create error page to show user
		}
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
}

