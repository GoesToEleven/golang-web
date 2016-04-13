# Put A File

We put a file with code like this:

```go
func putFile(ctx context.Context, name string, rdr io.Reader) error {

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	writer := client.Bucket(gcsBucket).Object(name).NewWriter(ctx)

	io.Copy(writer, rdr)
	// check for errors on io.Copy in production code!
	return writer.Close()
}
```

You can learn about the different types and methods in google/cloud by going to godoc.org and searching for `google/cloud` which will give you this link: [https://godoc.org/google.golang.org/cloud/storage](https://godoc.org/google.golang.org/cloud/storage)

My read through on the code would be like this:

In the `storage` package for Google cloud, you call `func newClient` which gives you a pointer to a client: `*storage.Client`.

Once you have a pointer to a client, there are two methods which you can call:

```go
func (c *Client) Bucket(name string) *BucketHandle
```

```go
func (c *Client) Close() error
```

Since there is a `Close()` method, this is a big hint that we should probably close a client. We can do this effectively with defer, as you see in the code sample above: `defer client.Close()`

To continue finding functionality in Google Cloud Storage, we will then also probably need to call `func (c *Client) Bucket(name string) *BucketHandle`. This makes sense as we will need to specify which bucket we want to access either to **put** or **get** a file.

So if we call `Bucket()` we are given a new type which is a pointer to a bucket handle: `*storage.BucketHandle`.

With a `*storage.BucketHandle`, we can see in the `storage` package documentation index that there are now several more methods available to us:
 
 ```go
 func (c *BucketHandle) ACL() *ACLHandle
 func (b *BucketHandle) Attrs(ctx context.Context) (*BucketAttrs, error)
 func (c *BucketHandle) DefaultObjectACL() *ACLHandle
 func (b *BucketHandle) List(ctx context.Context, q *Query) (*ObjectList, error)
 func (b *BucketHandle) Object(name string) *ObjectHandle
 ```
 
 

