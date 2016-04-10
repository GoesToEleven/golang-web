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

type demo struct {
	ctx    context.Context
	res    http.ResponseWriter
	bucket *storage.BucketHandle
	client *storage.Client
}

const gcsBucket = "learning-1130.appspot.com"

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

	d.createListFiles()
	d.listFiles()
	d.statFiles()
}

func (d *demo) statFiles() {
	io.WriteString(d.res, "\nRETRIEVING FILE STATS\n")

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

	for _, v := range objs.Results {
		d.statFile(v.Name)
	}
}

func (d *demo) statFile(fileName string) {
	io.WriteString(d.res, "\nFILE STAT:\n")

	obj, err := d.bucket.Object(fileName).Attrs(d.ctx)
	if err != nil {
		log.Errorf(d.ctx, "statFile: unable to stat file from bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}

	d.dumpStats(obj)
}

func (d *demo) dumpStats(obj *storage.ObjectAttrs) {
	fmt.Fprintf(d.res, "(filename: /%v/%v, ", obj.Bucket, obj.Name)
	fmt.Fprintf(d.res, "ContentType: %q, ", obj.ContentType)
	fmt.Fprintf(d.res, "ACL: %#v, ", obj.ACL)
	fmt.Fprintf(d.res, "Owner: %v, ", obj.Owner)
	fmt.Fprintf(d.res, "ContentEncoding: %q, ", obj.ContentEncoding)
	fmt.Fprintf(d.res, "Size: %v, ", obj.Size)
	fmt.Fprintf(d.res, "MD5: %q, ", obj.MD5)
	fmt.Fprintf(d.res, "CRC32C: %q, ", obj.CRC32C)
	fmt.Fprintf(d.res, "Metadata: %#v, ", obj.Metadata)
	fmt.Fprintf(d.res, "MediaLink: %q, ", obj.MediaLink)
	fmt.Fprintf(d.res, "StorageClass: %q, ", obj.StorageClass)
	if !obj.Deleted.IsZero() {
		fmt.Fprintf(d.res, "Deleted: %v, ", obj.Deleted)
	}
	fmt.Fprintf(d.res, "Updated: %v)\n", obj.Updated)
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

	for _, v := range objs.Results {
		io.WriteString(d.res, v.Name+"\n")
	}
}

func (d *demo) createListFiles() {
	io.WriteString(d.res, "\nCreating more files for listbucket...\n")
	for _, n := range []string{"foo1", "foo2", "bar", "bar/1", "bar/2", "boo/"} {
		d.createFile(n)
	}
}

func (d *demo) createFile(fileName string) {
	fmt.Fprintf(d.res, "Creating file /%v/%v\n", gcsBucket, fileName)

	wc := d.bucket.Object(fileName).NewWriter(d.ctx)
	wc.ContentType = "text/plain"
	wc.Metadata = map[string]string{
		"x-goog-meta-foo": "foo",
		"x-goog-meta-bar": "bar",
	}

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
