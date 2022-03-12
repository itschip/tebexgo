package tebexgo

import "log"

func (s *Session) GetAllPackages() (string, error) {
	resp, err := s.GetRequest("https://plugin.tebex.io/packages")
	if err != nil {
		log.Printf("Failed to fetch packages, Error: %v", err.Error())
		return "", nil
	}
	
	return string(resp), nil
}