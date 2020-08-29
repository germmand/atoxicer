package firestore

import (
	"cloud.google.com/go/firestore"
)

// Session wraps the Firestore client instance so that we can extend it and add methods to it.
type Session struct {
	FirestoreSession *firestore.Client
}
