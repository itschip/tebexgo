package tebexgo

import (
	"log"

	"github.com/itschip/tebexgo/internal"
)

func (s *Session) GetAllSales() {
	resp, err := internal.GetRequest(s.Secret, AllSalesEndpoint)
	if err != nil {
		log.Println("Failed to get all sales")
		return nil, err
	}

	// TODO: Unmarshal response to struct
}
