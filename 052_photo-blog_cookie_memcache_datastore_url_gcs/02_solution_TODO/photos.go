package mem

import (
	"crypto/sha1"
	"fmt"
	"google.golang.org/appengine"
	"io"
	"mime/multipart"
	"net/http"
"google.golang.org/cloud/storage"
"golang.org/x/net/context"
)

func uploadPhoto(mpf multipart.File, id string, req *http.Request) error {
	defer mpf.Close()
	fname := id + `/` + getSha(mpf) + `.jpg`
	return addPhoto(fname, id, req, mpf)
}

func addPhoto(fname string, id string, req *http.Request, rdr multipart.File) error {
	// all of the code below was added in this refactor
	// and previous "func addPhoto" code was deleted
	ctx := appengine.NewContext(req)

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	writer := client.Bucket(gcsBucket).Object(fname).NewWriter(ctx)
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	writer.ContentType = "image/jpeg"
	io.Copy(writer, rdr)
	return writer.Close()
}

func getSha(mpf multipart.File) string {
	h := sha1.New()
	io.Copy(h, mpf)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func listFiles(id string, ctx context.Context) ([]*storage.ObjectAttrs, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	q := storage.Query{
		Prefix: id,
	}

	ptr, err := client.Bucket(gcsBucket).List(ctx, &q)
	if err != nil {
		return nil, err
	}
	return ptr.Results, nil
}
