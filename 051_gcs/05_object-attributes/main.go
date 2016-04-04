package skyhdd

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"io"
	"net/http"
)

const gcsBucket = "learning-1130.appspot.com"

func init() {
	http.HandleFunc("/", handler)
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

		fnames, err := putCookie(res, req, fname)
		if err != nil {
			log.Errorf(ctx, "ERROR handler putCookie: ", err)
			http.Error(res, "We were unable to accept your file\n"+err.Error(), http.StatusUnsupportedMediaType)
			return
		}

		html += `<h1>Files</h1>`
		for k, _ := range fnames {
			attrs, err := getAttrs(ctx, k)
			if err != nil {
				continue
			}
			// TODO left off here
			html += `<h1>`+attrs.Name+`</h1>`+
			`<h3>Bucket: `+attrs.Bucket+`</h3>`+
			`<h3>ContentType: `+attrs.ContentType+`</h3>`+
			`<h3>ContentLanguage: `+attrs.ContentLanguage+`</h3>`+
			`<h3>CacheControl: `+attrs.CacheControl+`</h3>`+
			`<h3>ACL: `+attrs.ACL+`</h3>`+
			`<h3>`+attrs.Owner+`</h3>`+
			`<h3>`+attrs.Size+`</h3>`+
			`<h3>`+attrs.ContentEncoding+`</h3>`+
			`<h3>`+attrs.ContentDisposition+`</h3>`+
			`<h3>`+attrs.MD5+`</h3>`+
			`<h3>`+attrs.CRC32C+`</h3>`+
			`<h3>`+attrs.MediaLink+`</h3>`+
			`<h3>`+attrs.Metadata+`</h3>`+
			`<h3>`+attrs.Generation+`</h3>`+
			`<h3>`+attrs.MetaGeneration+`</h3>`+
			`<h3>`+attrs.StorageClass+`</h3>`+
			`<h3>`+attrs.Created+`</h3>`+
			`<h3>`+attrs.Deleted+`</h3>`+
			`<h3>`+attrs.Updated`</h3>`
		}
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, html)
}
