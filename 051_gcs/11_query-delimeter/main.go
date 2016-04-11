package skyhdd

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"io"
	"net/http"
	"strings"
)

func init() {
	http.HandleFunc("/", handler)
}

const gcsBucket = "learning-1130.appspot.com"

type demo struct {
	ctx    context.Context
	res    http.ResponseWriter
	bucket *storage.BucketHandle
	client *storage.Client
}

func handler(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	ctx := appengine.NewContext(req)

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "ERROR handler NewClient: ", err)
		return
	}
	defer client.Close()

	d := &demo{
		ctx:    ctx,
		res:    res,
		client: client,
		bucket: client.Bucket(gcsBucket),
	}

	d.createFiles()
	d.listFiles()
	io.WriteString(d.res, "\nRESULTS FROM LISTDIR - WITH DELIMITER\n")
	d.listDir()

}

func (d *demo) listDir() {

	query := &storage.Query{
		Delimiter: "/",
	}

	objs, err := d.bucket.List(d.ctx, query)
	if err != nil {
		log.Errorf(d.ctx, "listBucketDirMode: unable to list bucket %q: %v", gcsBucket, err)
		return
	}

	for _, obj := range objs.Results {
		fmt.Fprintf(d.res, "%v\n", obj.Name)
	}

	io.WriteString(d.res, "\nPREFIXES FOUND\n")

	fmt.Fprintf(d.res, "%v", objs.Prefixes)
}

func (d *demo) listFiles() {
	io.WriteString(d.res, "\nRETRIEVING FILE NAMES\n")

	client, err := storage.NewClient(d.ctx)
	if err != nil {
		log.Errorf(d.ctx, "%v", err)
		return
	}
	defer client.Close()

	objs, err := client.Bucket(gcsBucket).List(d.ctx, nil)
	if err != nil {
		log.Errorf(d.ctx, "%v", err)
		return
	}

	for _, obj := range objs.Results {
		io.WriteString(d.res, obj.Name+"\n")
	}
}

func (d *demo) createFiles() {
	io.WriteString(d.res, "\nCreating more files for listbucket...\n")
	for _, n := range []string{"foo1", "foo2", "bar", "bar/1", "bar/2", "boo/", "boo/yah", "compadre/amigo/diaz", "compadre/luego/hasta", "bar/nonce/1", "bar/nonce/2", "bar/nonce/compadre/1", "bar/nonce/compadre/2"} {
		d.createFile(n)
	}
}

func (d *demo) createFile(fileName string) {
	fmt.Fprintf(d.res, "Creating file /%v/%v\n", gcsBucket, fileName)

	wc := d.bucket.Object(fileName).NewWriter(d.ctx)
	wc.ContentType = "text/plain"

	if _, err := wc.Write([]byte("abcde\n")); err != nil {
		log.Errorf(d.ctx, "createFile: unable to write data to bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}
	if _, err := wc.Write([]byte(strings.Repeat("f", 1024*4) + "\n")); err != nil {
		log.Errorf(d.ctx, "createFile: unable to write data to bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}
	if err := wc.Close(); err != nil {
		log.Errorf(d.ctx, "createFile: unable to close bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}
}
