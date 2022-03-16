package tebexgo

import (
	"bytes"
	"encoding/json"
	"io"
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
	internal.UnmarshalResponse(res, &bans)

	return &bans, err
}

func (s *Session) CreateBan(input *BanInput) (*BanData, error) {
	jsonBody, err := json.Marshal(&input)
	bodyS := string(jsonBody)

	if err != nil {
		log.Println(err.Error())
	}

	reqBody := io.NopCloser(bytes.NewBufferString(bodyS))

	res, err := internal.PostRequest(s.Secret, BansEndpoint, reqBody)
	if err != nil {
		log.Println("Failed to create ban")
		return nil, err
	}

	var ban BanData
	internal.UnmarshalResponse(res, &ban)

	return &ban, nil
}
