package skyhdd

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"io"
	"net/http"
	"fmt"
)

const gcsBucket = "learning-1130.appspot.com"

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/golden", retriever)
}

func handler(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	html := `
		<h1>UPLOAD</h1>
	    <form method="POST" enctype="multipart/form-data">
		<input type="file" name="dahui">
		<input type="submit">
	    </form>
	`

	if req.Method == "POST" {

		mpf, hdr, err := req.FormFile("dahui")
		if err != nil {
			log.Errorf(ctx, "ERROR handler req.FormFile: ", err)
			http.Error(res, "We were unable to upload your file\n", http.StatusInternalServerError)
			return
		}
		defer mpf.Close()

		fname, err := uploadFile(req, mpf, hdr)
		if err != nil {
			log.Errorf(ctx, "ERROR handler uploadFile: ", err)
			http.Error(res, "We were unable to accept your file\n"+err.Error(), http.StatusUnsupportedMediaType)
			return
		}

		err = putCookie(res, req, fname)
		if err != nil {
			log.Errorf(ctx, "ERROR handler putCookie: ", err)
			http.Error(res, "We were unable to accept your file\n"+err.Error(), http.StatusUnsupportedMediaType)
			return
		}

	}

	fnames, err := getCookie(res, req)
	if err != nil {
		log.Infof(ctx, "INFO handler getCookie - no cookie yet: ", err)
	}

	html += `<h1>Files</h1>`

	if len(fnames) != 0 {
		for k, _ := range fnames {
			html += `<h3><a href="/golden?object=` + k + `">` + k + `</a></h3>`
		}
	}

	html += `<br>------------------------------<br>`

	xsAttrs, err := listFiles(ctx)
	if err != nil {
		log.Errorf(ctx, "ERROR handler listFiles: ", err)
		http.Error(res, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	for _, v := range xsAttrs {
		html += `<p><strong>Name:</strong><br> ` + v.Name + `</p>`+
		`<p><strong>Bucket:</strong><br> `+v.Bucket+`</p>` +
		`<p><strong>ContentType:</strong><br> `+v.ContentType+`</p>`+
		`<p><strong>ACL:</strong><br> `+fmt.Sprintf("%v",v.ACL)+`</p>`+
		`<p><strong>Owner:</strong><br>`+v.Owner+`</p>`+
		`<p><strong>MediaLink:</strong><br><a href="`+v.MediaLink+`" target="_blank">`+v.MediaLink+`</a></p>`
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, html)
}

func retriever(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	objectName := req.FormValue("object")
	rdr, err := getFile(ctx, objectName)
	if err != nil {
		log.Errorf(ctx, "ERROR golden getFile: ", err)
		http.Error(res, "We were unable to get the file"+objectName+"\n"+err.Error(), http.StatusUnsupportedMediaType)
		return
	}
	defer rdr.Close()
	io.Copy(res, rdr)
}
