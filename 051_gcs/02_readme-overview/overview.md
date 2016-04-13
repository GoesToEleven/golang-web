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
 
 ## ACL
 
 This will let us control our **Access Control List**. Basically these are settings which we can set on our **bucket** to control who can access the **bucket** and what they can do. This is known as *scope* and *permissions*.
  
   **Scope** defines who the permission applies to (for example, a specific user or group of users). Scopes are sometimes referred to as *grantees.* 

  **Permissions** define the actions that can be performed against a bucket (for example, read or write).
   
  We will also, later, be able to set ACL's for objects (files) which we store in Google Cloud Storage (GCS).
  
  More on this later.
 
 ## Attrs
 
 This will give us the attributes for a **bucket**. You can see the many different attributes at this link: [https://godoc.org/google.golang.org/cloud/storage#BucketAttrs](https://godoc.org/google.golang.org/cloud/storage#BucketAttrs) For convenience, I'm also listing them here:
 
 ```go
 type BucketAttrs struct {
     // Name is the name of the bucket.
     Name string
 
     // ACL is the list of access control rules on the bucket.
     ACL []ACLRule
 
     // DefaultObjectACL is the list of access controls to
     // apply to new objects when no object ACL is provided.
     DefaultObjectACL []ACLRule
 
     // Location is the location of the bucket. It defaults to "US".
     Location string
 
     // MetaGeneration is the metadata generation of the bucket.
     MetaGeneration int64
 
     // StorageClass is the storage class of the bucket. This defines
     // how objects in the bucket are stored and determines the SLA
     // and the cost of storage. Typical values are "STANDARD" and
     // "DURABLE_REDUCED_AVAILABILITY". Defaults to "STANDARD".
     StorageClass string
 
     // Created is the creation time of the bucket.
     Created time.Time
 }
 ```
 
  We will also, later, be able to see Attrs for objects (files) which we store in Google Cloud Storage (GCS).
  
  More on this later.