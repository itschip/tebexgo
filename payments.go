package tebexgo

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/itschip/tebexgo/internal"
)

// Retrive the latest payments (up to a maximum of 100) made on your webstore.
func (s *Session) GetAllPayments() (*[]Payment, error) {
	resp, err := internal.GetRequest(s.Secret, PaymentsEndpoint)
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

// Retrive a payment made on your webstore by transaction ID.
func (s *Session) RetrievePayment(transactionId string) (*Payment, error) {
	endpoint := fmt.Sprintf("%s/%s", PaymentsEndpoint, transactionId)
	resp, err := internal.GetRequest(s.Secret, endpoint)
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

// Returns an array of fields (custom variables, etc) required to
// be entered for a manual payment to be created for a package.
func (s *Session) GetRequiredPaymentFields(packageId string) (*[]PaymentPackageFields, error) {
	endpoint := fmt.Sprintf("%s/fields/%s", PaymentsEndpoint, packageId)
	resp, err := internal.GetRequest(s.Secret, endpoint)
	if err != nil {
		return nil, err
	}

	var paymentFields []PaymentPackageFields
	err = internal.UnmarshalResponse(resp, &paymentFields)
	if err != nil {
		return nil, err
	}

	return &paymentFields, nil
}

func (s *Session) UpdatePayment(transactionId string, object *UpdatePaymentObject) error {
	endpoint := fmt.Sprintf("%s/%s", PaymentsEndpoint, transactionId)
	jsonBody, err := json.Marshal(&object)
	if err != nil {
		return err
	}

	_, err = internal.PutRequest(s.Secret, endpoint, jsonBody)
	if err != nil {
		return err
	}

	return nil
}
