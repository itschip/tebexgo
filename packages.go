package tebexgo

import (
	"encoding/json"
	"github.com/itschip/tebexgo/internal"
	"log"
)

func (s *Session) GetAllPackages() (*[]Package, error) {
	resp, err := internal.GetRequest(s.Secret, AllPackagesEndpoint)
	
	if err != nil {
		log.Printf("Failed to fetch packages, Error: %v", err.Error())
		return nil, err
	}
	
	var packages []Package
	
	// Should abstract this into some function.
	err = json.Unmarshal(resp, &packages)
	if err != nil {
		log.Println("Failed to marshal response")
		return nil, err
	}
	
	return &packages, nil
}

func (s *Session) GetPackage(packageId string) (Package, error)  {
	resp, err := internal.GetRequest(s.Secret, RetrievePackageEndpoint + packageId)
	
	if err != nil {
		log.Printf("Failed to fetch packages, Error: %v", err.Error())
		return Package{}, err
	}
	
	var pkg Package
	
	// Same here
	err = json.Unmarshal(resp, &pkg)
	if err != nil {
		log.Println(err.Error())
		return Package{}, err
	}
	
	return pkg, nil
}