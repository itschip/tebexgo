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
	internal.UnmarshalResponse(res, &listing)
	
	return &listing, nil
}