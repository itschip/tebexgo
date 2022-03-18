package tebexgo

import (
	"fmt"
	"log"

	"github.com/itschip/tebexgo/internal"
)

func (s *Session) PlayerLookup(user string) (*PlayerData, error) {
	endpoint := fmt.Sprintf("%s/%s", PlayerLookupEndpoint, user)
	resp, err := internal.GetRequest(s.Secret, endpoint)
	if err != nil {
		log.Println("Failed to get player")
		return nil, err
	}

	var player PlayerData
	internal.UnmarshalResponse(resp, &player)

	return &player, nil
}
