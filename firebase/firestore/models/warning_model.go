package models

// Warning represents the amount of warnings a user has received within a guild/server.
// 3 Red warnings and the user should be muted.
// 2 Yello warnings represents 1 red warning.
type Warning struct {
	UserID         string `firestore:"userid,omitempty"`
	GuildID        string `firestore:"guildid,omitempty"`
	RedWarnings    int    `firestore:"redwarnings,omitempty"`
	YellowWarnings int    `firestore:"yellowwarnings,omitempty"`
}
