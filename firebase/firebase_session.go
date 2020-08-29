package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebasego "firebase.google.com/go"
)

type FirebaseSession struct {
	FirebaseApp *firebasego.App
}

func NewApp(ctx context.Context) *FirebaseSession {
	app, err := firebasego.NewApp(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	return &FirebaseSession{
		FirebaseApp: app,
	}
}

func (app *FirebaseSession) NewFirestoreSession(ctx context.Context) *firestore.Client {
	firestoreSession, err := app.FirebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return firestoreSession
}
