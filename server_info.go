package tebexgo

import (
	"github.com/itschip/tebexgo/internal"
	"log"
)

func (s *Session) GetServerInfo() (*ServerInformation, error) {
	resp, err := internal.GetRequest(s.Secret, AllPackagesEndpoint)
	
	if err != nil {
		log.Printf("Failed to fetch server information, Error: %v", err.Error())
		return nil, err
	}
	
	var serverInfo ServerInformation
	
	internal.UnmarshalResponse(resp, &serverInfo)
	
	return &serverInfo, nil
}