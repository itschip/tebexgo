package tebexgo

import (
	"github.com/itschip/tebexgo/internal"
	"log"
)

func (s *Session) GetListing() (*Listing, error) {
	res, err := internal.GetRequest(s.Secret, ListingEndpoint)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	
	var listing Listing
	if err := internal.UnmarshalResponse(res, &listing); err != nil {
		return nil, err
	}
	
	return &listing, nil
}