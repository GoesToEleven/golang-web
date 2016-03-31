package gcsdemo

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
)

// bucket is a local cache of the app's default bucket name.
var bucket = "learning-1130.appspot.com"

func init() {
	http.HandleFunc("/", handler)
}

// demo struct holds information needed to run the various demo functions.
type demo struct {
	bucket *storage.BucketHandle
	client *storage.Client

	res http.ResponseWriter
	ctx context.Context
	// cleanUp is a list of filenames that need cleaning up at the end of the demo.
	cleanUp []string
	// failed indicates that one or more of the demo steps failed.
	failed bool
}

func (d *demo) errorf(format string, args ...interface{}) {
	d.failed = true
	log.Errorf(d.ctx, format, args...)
}

// handler is the main demo entry point that calls the GCS operations.
func handler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	ctx := appengine.NewContext(req)
	if bucket == "" {
		var err error
		if bucket, err = file.DefaultBucketName(ctx); err != nil {
			log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
			return
		}
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
		return
	}
	defer client.Close()

	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(res, "Demo GCS Application running from Version: %v\n", appengine.VersionID(ctx))
	fmt.Fprintf(res, "Using bucket name: %v\n\n", bucket)

	d := &demo{
		res:    res,
		ctx:    ctx,
		client: client,
		bucket: client.Bucket(bucket),
	}

	n := "demo-testfile-go"
	d.createFile(n)
	d.readFile(n)
	d.copyFile(n)
	d.statFile(n)
	d.createListFiles()
	d.listBucket()
	d.listBucketDirMode()
	d.defaultACL()
	d.putDefaultACLRule()
	d.deleteDefaultACLRule()
	d.bucketACL()
	d.putBucketACLRule()
	d.deleteBucketACLRule()
	d.acl(n)
	d.putACLRule(n)
	d.deleteACLRule(n)
	d.deleteFiles()

	if d.failed {
		io.WriteString(res, "\nDemo failed.\n")
	} else {
		io.WriteString(res, "\nDemo succeeded.\n")
	}
}

//[START write]
// createFile creates a file in Google Cloud Storage.
func (d *demo) createFile(fileName string) {
	fmt.Fprintf(d.res, "Creating file /%v/%v\n", bucket, fileName)

	wc := d.bucket.Object(fileName).NewWriter(d.ctx)
	wc.ContentType = "text/plain"
	wc.Metadata = map[string]string{
		"x-goog-meta-foo": "foo",
		"x-goog-meta-bar": "bar",
	}
	d.cleanUp = append(d.cleanUp, fileName)

	if _, err := wc.Write([]byte("abcde\n")); err != nil {
		d.errorf("createFile: unable to write data to bucket %q, file %q: %v", bucket, fileName, err)
		return
	}
	if _, err := wc.Write([]byte(strings.Repeat("f", 1024*4) + "\n")); err != nil {
		d.errorf("createFile: unable to write data to bucket %q, file %q: %v", bucket, fileName, err)
		return
	}
	if err := wc.Close(); err != nil {
		d.errorf("createFile: unable to close bucket %q, file %q: %v", bucket, fileName, err)
		return
	}
}

//[END write]

//[START read]
// readFile reads the named file in Google Cloud Storage.
func (d *demo) readFile(fileName string) {
	io.WriteString(d.res, "\nAbbreviated file content (first line and last 1K):\n")

	rc, err := d.bucket.Object(fileName).NewReader(d.ctx)
	if err != nil {
		d.errorf("readFile: unable to open file from bucket %q, file %q: %v", bucket, fileName, err)
		return
	}
	defer rc.Close()
	slurp, err := ioutil.ReadAll(rc)
	if err != nil {
		d.errorf("readFile: unable to read data from bucket %q, file %q: %v", bucket, fileName, err)
		return
	}

	fmt.Fprintf(d.res, "%s\n", bytes.SplitN(slurp, []byte("\n"), 2)[0])
	if len(slurp) > 1024 {
		fmt.Fprintf(d.res, "...%s\n", slurp[len(slurp)-1024:])
	} else {
		fmt.Fprintf(d.res, "%s\n", slurp)
	}
}

//[END read]

// copyFile copies a file in Google Cloud Storage.
func (d *demo) copyFile(fileName string) {
	copyName := fileName + "-copy"
	fmt.Fprintf(d.res, "Copying file /%v/%v to /%v/%v:\n", bucket, fileName, bucket, copyName)

	obj, err := d.bucket.Object(fileName).CopyTo(d.ctx, d.bucket.Object(copyName), nil)
	if err != nil {
		d.errorf("copyFile: unable to copy /%v/%v to bucket %q, file %q: %v", bucket, fileName, bucket, copyName, err)
		return
	}
	d.cleanUp = append(d.cleanUp, copyName)

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

// statFile reads the stats of the named file in Google Cloud Storage.
func (d *demo) statFile(fileName string) {
	io.WriteString(d.res, "\nFile stat:\n")

	obj, err := d.bucket.Object(fileName).Attrs(d.ctx)
	if err != nil {
		d.errorf("statFile: unable to stat file from bucket %q, file %q: %v", bucket, fileName, err)
		return
	}

	d.dumpStats(obj)
}

// createListFiles creates files that will be used by listBucket.
func (d *demo) createListFiles() {
	io.WriteString(d.res, "\nCreating more files for listbucket...\n")
	for _, n := range []string{"foo1", "foo2", "bar", "bar/1", "bar/2", "boo/"} {
		d.createFile(n)
	}
}

// listBucket lists the contents of a bucket in Google Cloud Storage.
func (d *demo) listBucket() {
	io.WriteString(d.res, "\nListbucket result:\n")

	query := &storage.Query{Prefix: "foo"}
	for query != nil {
		objs, err := d.bucket.List(d.ctx, query)
		if err != nil {
			d.errorf("listBucket: unable to list bucket %q: %v", bucket, err)
			return
		}
		query = objs.Next

		for _, obj := range objs.Results {
			d.dumpStats(obj)
		}
	}
}

func (d *demo) listDir(name, indent string) {
	query := &storage.Query{Prefix: name, Delimiter: "/"}
	for query != nil {
		objs, err := d.bucket.List(d.ctx, query)
		if err != nil {
			d.errorf("listBucketDirMode: unable to list bucket %q: %v", bucket, err)
			return
		}
		query = objs.Next

		for _, obj := range objs.Results {
			fmt.Fprint(d.res, indent)
			d.dumpStats(obj)
		}
		for _, dir := range objs.Prefixes {
			fmt.Fprintf(d.res, "%v(directory: /%v/%v)\n", indent, bucket, dir)
			d.listDir(dir, indent+"  ")
		}
	}
}

// listBucketDirMode lists the contents of a bucket in dir mode in Google Cloud Storage.
func (d *demo) listBucketDirMode() {
	io.WriteString(d.res, "\nListbucket directory mode result:\n")
	d.listDir("b", "")
}

// dumpDefaultACL prints out the default object ACL for this bucket.
func (d *demo) dumpDefaultACL() {
	acl, err := d.bucket.ACL().List(d.ctx)
	if err != nil {
		d.errorf("defaultACL: unable to list default object ACL for bucket %q: %v", bucket, err)
		return
	}
	for _, v := range acl {
		fmt.Fprintf(d.res, "Scope: %q, Permission: %q\n", v.Entity, v.Role)
	}
}

// defaultACL displays the default object ACL for this bucket.
func (d *demo) defaultACL() {
	io.WriteString(d.res, "\nDefault object ACL:\n")
	d.dumpDefaultACL()
}

// putDefaultACLRule adds the "allUsers" default object ACL rule for this bucket.
func (d *demo) putDefaultACLRule() {
	io.WriteString(d.res, "\nPut Default object ACL Rule:\n")
	err := d.bucket.DefaultObjectACL().Set(d.ctx, storage.AllUsers, storage.RoleReader)
	if err != nil {
		d.errorf("putDefaultACLRule: unable to save default object ACL rule for bucket %q: %v", bucket, err)
		return
	}
	d.dumpDefaultACL()
}

// deleteDefaultACLRule deleted the "allUsers" default object ACL rule for this bucket.
func (d *demo) deleteDefaultACLRule() {
	io.WriteString(d.res, "\nDelete Default object ACL Rule:\n")
	err := d.bucket.DefaultObjectACL().Delete(d.ctx, storage.AllUsers)
	if err != nil {
		d.errorf("deleteDefaultACLRule: unable to delete default object ACL rule for bucket %q: %v", bucket, err)
		return
	}
	d.dumpDefaultACL()
}

// dumpBucketACL prints out the bucket ACL.
func (d *demo) dumpBucketACL() {
	acl, err := d.bucket.ACL().List(d.ctx)
	if err != nil {
		d.errorf("dumpBucketACL: unable to list bucket ACL for bucket %q: %v", bucket, err)
		return
	}
	for _, v := range acl {
		fmt.Fprintf(d.res, "Scope: %q, Permission: %q\n", v.Entity, v.Role)
	}
}

// bucketACL displays the bucket ACL for this bucket.
func (d *demo) bucketACL() {
	io.WriteString(d.res, "\nBucket ACL:\n")
	d.dumpBucketACL()
}

// putBucketACLRule adds the "allUsers" bucket ACL rule for this bucket.
func (d *demo) putBucketACLRule() {
	io.WriteString(d.res, "\nPut Bucket ACL Rule:\n")
	err := d.bucket.ACL().Set(d.ctx, storage.AllUsers, storage.RoleReader)
	if err != nil {
		d.errorf("putBucketACLRule: unable to save bucket ACL rule for bucket %q: %v", bucket, err)
		return
	}
	d.dumpBucketACL()
}

// deleteBucketACLRule deleted the "allUsers" bucket ACL rule for this bucket.
func (d *demo) deleteBucketACLRule() {
	io.WriteString(d.res, "\nDelete Bucket ACL Rule:\n")
	err := d.bucket.ACL().Delete(d.ctx, storage.AllUsers)
	if err != nil {
		d.errorf("deleteBucketACLRule: unable to delete bucket ACL rule for bucket %q: %v", bucket, err)
		return
	}
	d.dumpBucketACL()
}

// dumpACL prints out the ACL of the named file.
func (d *demo) dumpACL(fileName string) {
	acl, err := d.bucket.Object(fileName).ACL().List(d.ctx)
	if err != nil {
		d.errorf("dumpACL: unable to list file ACL for bucket %q, file %q: %v", bucket, fileName, err)
		return
	}
	for _, v := range acl {
		fmt.Fprintf(d.res, "Scope: %q, Permission: %q\n", v.Entity, v.Role)
	}
}

// acl displays the ACL for the named file.
func (d *demo) acl(fileName string) {
	fmt.Fprintf(d.res, "\nACL for file %v:\n", fileName)
	d.dumpACL(fileName)
}

// putACLRule adds the "allUsers" ACL rule for the named file.
func (d *demo) putACLRule(fileName string) {
	fmt.Fprintf(d.res, "\nPut ACL rule for file %v:\n", fileName)
	err := d.bucket.Object(fileName).ACL().Set(d.ctx, storage.AllUsers, storage.RoleReader)
	if err != nil {
		d.errorf("putACLRule: unable to save ACL rule for bucket %q, file %q: %v", bucket, fileName, err)
		return
	}
	d.dumpACL(fileName)
}

// deleteACLRule deleted the "allUsers" ACL rule for the named file.
func (d *demo) deleteACLRule(fileName string) {
	fmt.Fprintf(d.res, "\nDelete ACL rule for file %v:\n", fileName)
	err := d.bucket.Object(fileName).ACL().Delete(d.ctx, storage.AllUsers)
	if err != nil {
		d.errorf("deleteACLRule: unable to delete ACL rule for bucket %q, file %q: %v", bucket, fileName, err)
		return
	}
	d.dumpACL(fileName)
}

// deleteFiles deletes all the temporary files from a bucket created by this demo.
func (d *demo) deleteFiles() {
	io.WriteString(d.res, "\nDeleting files...\n")
	for _, v := range d.cleanUp {
		fmt.Fprintf(d.res, "Deleting file %v\n", v)
		if err := d.bucket.Object(v).Delete(d.ctx); err != nil {
			d.errorf("deleteFiles: unable to delete bucket %q, file %q: %v", bucket, v, err)
			return
		}
	}
}

//[END sample]
