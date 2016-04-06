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
		{storage.ACLEntity("tscottmcleod@gmail.com"), storage.RoleReader},
	}
	io.Copy(writer, rdr)
	return writer.Close()
}

func getFile(ctx context.Context, name string) (io.ReadCloser, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return client.Bucket(gcsBucket).Object(name).NewReader(ctx)
}

func listFiles(ctx context.Context) ([]*storage.ObjectAttrs, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	q := storage.Query{
		Versions: false,
		MaxResults: 2,
	}

	ptr, err := client.Bucket(gcsBucket).List(ctx, &q)
	if err != nil {
		return nil, err
	}
	return ptr.Results, nil
}