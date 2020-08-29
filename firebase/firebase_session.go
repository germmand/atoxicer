package firebase

import (
	"context"
	"log"

	firebasego "firebase.google.com/go"
	"github.com/germmand/atoxicer/firebase/firestore"
)

// Session wraps the Firebase App so that we can extend it and add methods to it.
type Session struct {
	FirebaseApp *firebasego.App
}

// NewApp creates a new instance of a Session.
// This is because Session should no be created manually.
func NewApp(ctx context.Context) *Session {
	app, err := firebasego.NewApp(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	return &Session{
		FirebaseApp: app,
	}
}

// NewFirestoreSession extracts the Firestore session from the Firebase session.
// TODO: Remove these unnecesary comments. This is to shut the fucking go-lint up.
func (app *Session) NewFirestoreSession(ctx context.Context) *firestore.Session {
	firestoreSession, err := app.FirebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return &firestore.Session{
		FirestoreSession: firestoreSession,
	}
}
