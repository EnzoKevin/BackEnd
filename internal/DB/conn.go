package db

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type FirebaseDB struct {
	App       *firebase.App
	Auth      *auth.Client
	Firestore *firestore.Client
}

func ConnectDB() (*FirebaseDB, error) {
	ctx := context.Background()

	credPath := os.Getenv("FIREBASE_CREDENTIALS")
	if credPath == "" {
		return nil, fmt.Errorf("variável FIREBASE_CREDENTIALS não definida")
	}

	opt := option.WithCredentialsFile(credPath)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("🔥 Connected to Firebase")

	return &FirebaseDB{
		App:       app,
		Auth:      authClient,
		Firestore: firestoreClient,
	}, nil
}
