package tebexgo

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/itschip/tebexgo/internal"
)

func (s *Session) GetAllBans() (*Bans, error) {
	res, err := internal.GetRequest(s.Secret, BansEndpoint)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var bans Bans
	err = internal.UnmarshalResponse(res, &bans)
	if err != nil {
		return nil, err
	}

	return &bans, err
}

func (s *Session) CreateBan(input *BanInput) (*Bans, error) {
	jsonBody, err := json.Marshal(&input)
	if err != nil {
		return nil, err
	}

	res, err := internal.PostRequest(s.Secret, BansEndpoint, jsonBody)
	if err != nil {
		log.Println("Failed to create ban")
		return nil, err
	}

	fmt.Println(string(res))

	var ban Bans
	internal.UnmarshalResponse(res, &ban)

	return &ban, nil
}
