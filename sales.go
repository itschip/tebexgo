package tebexgo

import (
	"log"

	"github.com/itschip/tebexgo/internal"
)

func (s *Session) GetAllSales() (*Sales, error) {
	resp, err := internal.GetRequest(s.Secret, AllSalesEndpoint)
	if err != nil {
		log.Println("Failed to get all sales")
		return nil, err
	}
	
	var sales Sales
	err = internal.UnmarshalResponse(resp, &sales)
	if err != nil {
		return nil, err
	}
	
	return &sales, nil
}
