package tebexgo

import (
	"encoding/json"
	"fmt"
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

	if err := internal.UnmarshalResponse(resp, &packages); err != nil {
		return nil, err
	}

	return &packages, nil
}

func (s *Session) GetPackage(packageId string) (*Package, error) {
	resp, err := internal.GetRequest(s.Secret, RetrievePackageEndpoint+packageId)

	if err != nil {
		log.Printf("Failed to fetch package, Error: %v", err.Error())
		return nil, err
	}

	var pkg Package

	if err := internal.UnmarshalResponse(resp, &pkg); err != nil {
		return nil, err
	}

	return &pkg, nil
}

func (s *Session) UpdatePackage(packageId string, updateObject *UpdatePackageObject) error {
	endpoint := fmt.Sprintf("%s/%s", RetrievePackageEndpoint, packageId)

	reqBody, err := json.Marshal(&updateObject)
	if err != nil {
		log.Println("Failed to marshal body")
		return err
	}

	_, err = internal.PutRequest(s.Secret, endpoint, reqBody)

	if err != nil {
		log.Printf("Failed to update package, Error: %v", err.Error())
		return err
	}

	return nil
}
