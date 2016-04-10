package skyhdd

import (
	"golang.org/x/net/context"
	"google.golang.org/cloud/storage"
)

func listFiles(ctx context.Context) (*storage.ObjectList, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	objs, err := client.Bucket(gcsBucket).List(ctx, nil)
	if err != nil {
		return nil, err
	}
	return objs, nil
}