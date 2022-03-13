package tebexgo

import (
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
	
	internal.UnmarshalResponse(resp, &packages)
	
	return &packages, nil
}

func (s *Session) GetPackage(packageId string) (*Package, error)  {
	resp, err := internal.GetRequest(s.Secret, RetrievePackageEndpoint + packageId)
	
	if err != nil {
		log.Printf("Failed to fetch package, Error: %v", err.Error())
		return nil, err
	}
	
	var pkg Package
	
	internal.UnmarshalResponse(resp, &pkg)
	
	return &pkg, nil
}