package skyhdd

import (
	"golang.org/x/net/context"
	"google.golang.org/cloud/storage"
	"io"
)

func putFile(ctx context.Context, name string, rdr io.Reader) error {

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	writer := client.Bucket(gcsBucket).Object(name).NewWriter(ctx)
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	io.Copy(writer, rdr)
	return writer.Close()
}

// getAttrs returns a *storage.ObjectAttrs - learn more about this Type here:
// https://godoc.org/google.golang.org/cloud/storage#ObjectAttrs
func getAttrs(ctx context.Context, name string) (*storage.ObjectAttrs, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	attrs, err := client.Bucket(gcsBucket).Object(name).Attrs(ctx)
	if err != nil {
		return nil, err
	}
	return attrs, nil
}
