package tebexgo

import (
	"fmt"
	"log"

	"github.com/itschip/tebexgo/internal"
)

func (s *Session) GetAllPayments() (*[]Payment, error) {
	resp, err := internal.GetRequest(s.Secret, PaymentsEndpoint, nil)
	if err != nil {
		log.Println("Failed to get all payments")
		return nil, err
	}

	var payments []Payment
	err = internal.UnmarshalResponse(resp, &payments)
	if err != nil {
		log.Println("Failed to unmarshal response")
		return nil, err
	}

	return &payments, nil
}

func (s *Session) RetrievePayment(transactionId string) (*Payment, error) {
	endpoint := fmt.Sprintf("%s/%s", PaymentsEndpoint, transactionId)
	resp, err := internal.GetRequest(s.Secret, endpoint, nil)
	if err != nil {
		log.Println("Failed to get payment")
		return nil, err
	}

	var payment Payment
	err = internal.UnmarshalResponse(resp, &payment)
	if err != nil {
		log.Println("Failed to unmarshal response")
		return nil, err
	}

	return &payment, err
}
