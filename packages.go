package tebexgo

import (
	"encoding/json"
	"log"

	"github.com/itschip/tebexgo/internal"
)

func (s *Session) GetAllPackages() (*[]Package, error) {
	resp, err := internal.GetRequest(s.Secret, AllPackagesEndpoint)

	if err != nil {
		log.Printf("Failed to fetch packages, Error: %v", err.Error())
		return nil, err
	}

	var packages []Package

	internal.UnmarshalResponse(resp, &packages)

	return &packages, nil
}

func (s *Session) GetPackage(packageId string) (*Package, error) {
	resp, err := internal.GetRequest(s.Secret, RetrievePackageEndpoint+packageId)

	if err != nil {
		log.Printf("Failed to fetch package, Error: %v", err.Error())
		return nil, err
	}

	var pkg Package

	internal.UnmarshalResponse(resp, &pkg)

	return &pkg, nil
}

func (s *Session) UpdatePackage(packageId string, updateObject *UpdatePackageObject) error {
	reqBody, err := json.Marshal(&updateObject)
	if err != nil {
		log.Println("Failed to marshal body")
		return err
	}

	_, err = internal.PutRequest(s.Secret, RetrievePackageEndpoint+packageId, reqBody)

	if err != nil {
		log.Printf("Failed to update package, Error: %v", err.Error())
		return err
	}

	return nil
}
