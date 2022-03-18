package tebexgo

import (
	"github.com/itschip/tebexgo/internal"
	"log"
)

func (s *Session) GetServerInfo() (*ServerInformation, error) {
	resp, err := internal.GetRequest(s.Secret, InformationEndpoint)
	
	if err != nil {
		log.Printf("Failed to fetch server information, Error: %v", err.Error())
		return nil, err
	}
	
	var serverInfo ServerInformation
	
	if err := internal.UnmarshalResponse(resp, &serverInfo); err != nil {
		return nil, err
	}
	
	return &serverInfo, nil
}