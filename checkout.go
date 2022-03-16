package tebexgo

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/itschip/tebexgo/internal"
)

func (s *Session) CreateCheckoutUrl(checkoutObject *PutCheckoutObject) (*Checkout, error) {
	fmt.Println(checkoutObject)
	reqBody, err := json.Marshal(&checkoutObject)
	if err != nil {
		log.Println("Failed to marshal checkout body")
	}

	fmt.Println(string(reqBody))

	res, err := internal.PostRequest(s.Secret, CheckoutEndpoint, reqBody)
	if err != nil {
		log.Println("Failed to create checkout url")
		return nil, err
	}

	var checkout Checkout
	err = internal.UnmarshalResponse(res, &checkout)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(checkout)

	return &checkout, nil
}
