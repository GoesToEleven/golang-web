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
			html += `<h1>`+attrs.Name+`</h1>`
			//`<p>Bucket: `+attrs.Bucket+`</p>`+
			//`<p>ContentType: `+attrs.ContentType+`</p>`+
			//`<p>ContentLanguage: `+attrs.ContentLanguage+`</p>`+
			//`<p>CacheControl: `+attrs.CacheControl+`</p>`+
			////`<p>ACL: `+attrs.ACL+`</p>`+
			//`<p>Owner:`+attrs.Owner+`</p>`+
			////`<p>Size:`+attrs.Size+`</p>`+
			//`<p>ContentEncoding:`+attrs.ContentEncoding+`</p>`+
			//`<p>ContentDisposition:`+attrs.ContentDisposition+`</p>`+
			////`<p>MD5:`+attrs.MD5+`</p>`+
			//`<p>CRC32C:`+attrs.CRC32C+`</p>`+
			//`<p>MediaLink:`+attrs.MediaLink+`</p>`+
			//`<p>Metadata:`+attrs.Metadata+`</p>`+
			//`<p>Generation:`+attrs.Generation+`</p>`+
			//`<p>MetaGeneration:`+attrs.MetaGeneration+`</p>`+
			//`<p>StorageClass:`+attrs.StorageClass+`</p>`+
			//`<p>Created:`+attrs.Created+`</p>`+
			//`<p>Deleted:`+attrs.Deleted+`</p>`+
			//`<p>Updated:`+attrs.Updated+`</p>`
		}
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, html)
}
