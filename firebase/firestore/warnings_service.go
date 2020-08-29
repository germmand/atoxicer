package firestore

import (
	"context"

	"github.com/germmand/atoxicer/firebase/firestore/models"
	"google.golang.org/api/iterator"
)

// RetrieveWarning retrieves a warning
// Fuck I need to the remove these unnecessary comments. (This is due to go-lint)
func (s *Session) RetrieveWarning(ctx context.Context, userID string, guildID string) (*models.Warning, error) {
	var warningUser models.Warning
	warningCollection := s.FirestoreSession.Collection("warnings")
	iter := warningCollection.Where("userid", "==", userID).Where("guildid", "==", guildID).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		doc.DataTo(&warningUser)
		break
	}
	return &warningUser, nil
}

// SetWarning either saves or updates a warning...
func (s *Session) SetWarning(ctx context.Context, docID string, warning *models.Warning) error {
	warningCollection := s.FirestoreSession.Collection("warnings")
	_, err := warningCollection.Doc(docID).Set(ctx, warning)
	return err
}

// DeleteWarning deletes a warning duh...
// For real tho, I need to shut go-lint up. These fucking useless comments are annoying.
func (s *Session) DeleteWarning(ctx context.Context, docID string) error {
	warningCollection := s.FirestoreSession.Collection("warnings")
	_, err := warningCollection.Doc(docID).Delete(ctx)
	return err
}
