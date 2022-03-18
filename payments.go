package tebexgo

import (
	"log"

	"github.com/itschip/tebexgo/internal"
)

func (s *Session) GetAllPayments() ([]Payment, error) {
	resp, err := internal.GetRequest()
	if err != nil {
		log.Println("Failed to get all payments")
		return nil, err
	}

	var payments []Payment
	err := internal.UnmarshalResponse(resp, &payments)
	if err != nil {
		log.Println("Failed to unmarshal response")
		return nil, err
	}

	return &payments, nil
}
