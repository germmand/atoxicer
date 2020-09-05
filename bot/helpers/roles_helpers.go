package helpers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func FilterRoleByName(roles []*discordgo.Role, roleName string) (string, error) {
	for _, r := range roles {
		if r.Name == roleName {
			return r.ID, nil
		}
	}
	return "", fmt.Errorf("no role named: %s", roleName)
}
