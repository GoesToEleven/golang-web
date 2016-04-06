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

		_, err = putCookie(res, req, fname)
		if err != nil {
			log.Errorf(ctx, "ERROR handler putCookie: ", err)
			http.Error(res, "We were unable to accept your file\n"+err.Error(), http.StatusUnsupportedMediaType)
			return
		}
	}

	html += `<h1>Files</h1>`

	xsAttrs, err := listFiles(ctx)
	if err != nil {
		log.Errorf(ctx, "ERROR handler listFiles: ", err)
		http.Error(res, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	for _, v := range xsAttrs {
		html += `<h3>` + v.Name + `</h3>`+
		`<p><strong>Bucket:</strong><br> `+v.Bucket+`</p>` +
		`<p><strong>ContentType:</strong><br> `+v.ContentType+`</p>`+
		`<p><strong>ContentLanguage:</strong><br> `+v.ContentLanguage+`</p>`+
		`<p><strong>CacheControl:</strong><br> `+v.CacheControl+`</p>`+
		`<p><strong>ACL:</strong><br> `+fmt.Sprintf("%v",v.ACL)+`</p>`+
		`<p><strong>Owner:</strong><br>`+v.Owner+`</p>`+
		`<p><strong>Size:</strong><br>`+fmt.Sprintf("%v",v.Size)+`</p>`+
		`<p><strong>ContentEncoding:</strong><br>`+v.ContentEncoding+`</p>`+
		`<p><strong>ContentDisposition:</strong><br>`+v.ContentDisposition+`</p>`+
		`<p><strong>MD5:</strong><br>`+fmt.Sprintf("%v",v.MD5)+`</p>`+
		`<p><strong>CRC32C:</strong><br>`+fmt.Sprintf("%v",v.CRC32C)+`</p>`+
		`<p><strong>MediaLink:</strong><br><a href="`+v.MediaLink+`" target="_blank">`+v.MediaLink+`</a></p>`+
		`<p><strong>Metadata:</strong><br>`+fmt.Sprintf("%v",v.Metadata)+`</p>`+
		`<p><strong>Generation:</strong><br>`+fmt.Sprintf("%v",v.Generation)+`</p>`+
		`<p><strong>MetaGeneration:</strong><br>`+fmt.Sprintf("%v",v.MetaGeneration)+`</p>`+
		`<p><strong>StorageClass:</strong><br>`+v.StorageClass+`</p>`+
		`<p><strong>Created:</strong><br>`+fmt.Sprintf("%v",v.Created)+`</p>`+
		`<p><strong>Deleted:</strong><br>`+fmt.Sprintf("%v",v.Deleted)+`</p>`+
		`<p><strong>Updated:</strong><br>`+fmt.Sprintf("%v",v.Updated)+`</p>`
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, html)
}