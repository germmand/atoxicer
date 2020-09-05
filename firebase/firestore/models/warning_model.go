package models

import (
	"github.com/germmand/atoxicer/bot/constants"
)

// Warning represents the amount of warnings a user has received within a guild/server.
// 3 Red warnings and the user should be muted.
// 2 Yello warnings represents 1 red warning.
type Warning struct {
	UserID         string `firestore:"userid,omitempty"`
	GuildID        string `firestore:"guildid,omitempty"`
	RedWarnings    int    `firestore:"redwarnings,omitempty"`
	YellowWarnings int    `firestore:"yellowwarnings,omitempty"`
}

// UpdateWarningUponToxicity updates warning duh..
// TODO: Remove annoying warnings from go-lint.
func (w *Warning) UpdateWarningUponToxicity(toxicity constants.ToxicType) {
	if toxicity == constants.ToxicTypeHigh {
		w.RedWarnings++
	} else if toxicity == constants.ToxicTypeMedium {
		w.YellowWarnings++
	}

	if w.YellowWarnings >= 2 {
		w.YellowWarnings = 0
		w.RedWarnings++
	}
}
