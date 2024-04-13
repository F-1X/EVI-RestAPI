package database

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type DB interface {
	Collection(path string) *firestore.CollectionRef
	Close(ctx context.Context) error
}

type firestoreClientWrapper struct {
	client *firestore.Client
}

func (w *firestoreClientWrapper) Collection(path string) *firestore.CollectionRef {

	return w.client.Collection(path)
}

func InitFirestoreClient(ctx context.Context, credsFilePath string) (DB, error) {
	client, err := initFirestoreClient(ctx, credsFilePath)
	if err != nil {
		return nil, err
	}

	return &firestoreClientWrapper{client: client}, nil
}

func initFirestoreClient(ctx context.Context, credsFilePath string) (*firestore.Client, error) {
	sa := option.WithCredentialsFile(credsFilePath)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, err
	}

	return app.Firestore(ctx)
}

func (w *firestoreClientWrapper) Close(ctx context.Context) error {
	return w.client.Close()
}
