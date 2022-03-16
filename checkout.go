package tebexgo

import (
	"bytes"
	"encoding/json"
	"io"
	"log"

	"github.com/itschip/tebexgo/internal"
)

func (s *Session) CreateCheckoutUrl(checkoutObject *PutCheckoutObject) (*Checkout, error) {
	reqBody, err := json.Marshal(&checkoutObject)
	if err != nil {
		log.Println("Failed to marshal checkout body")
	}

	bodyS := string(reqBody)
	jsonBody := io.NopCloser(bytes.NewBufferString(bodyS))

	res, err := internal.PostRequest(s.Secret, CheckoutEndpoint, jsonBody)
	if err != nil {
		log.Println("Failed to create checkout url")
		return nil, err
	}

	var checkout Checkout
	internal.UnmarshalResponse(res, &checkout)

	return &checkout, nil
}
