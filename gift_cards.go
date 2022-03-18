package tebexgo

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/itschip/tebexgo/internal"
)

func (s *Session) GetAllGiftCards() (*GiftCards, error) {
	resp, err := internal.GetRequest(s.Secret, GiftCardsEndpoint)
	if err != nil {
		log.Println("Failed to get all gift cards")
		return nil, err
	}

	var giftCards GiftCards
	internal.UnmarshalResponse(resp, &giftCards)

	return &giftCards, nil
}

func (s *Session) GetGiftCard(cardId string) (*GiftCard, error) {
	endpoint := fmt.Sprintf("%s/%s", GiftCardsEndpoint, cardId)
	resp, err := internal.GetRequest(s.Secret, endpoint)
	if err != nil {
		log.Println("Failed to get gift card")
		return nil, err
	}

	var giftCard GiftCard
	internal.UnmarshalResponse(resp, &giftCard)

	return &giftCard, nil
}

func (s *Session) CreateGiftCard(cardObject *PutGiftCardObject) (*GiftCard, error) {
	jsonBody, err := json.Marshal(&cardObject)
	if err != nil {
		log.Println("Failed to marshal gift card object")
		return nil, err
	}

	resp, err := internal.PostRequest(s.Secret, GiftCardsEndpoint, jsonBody)
	if err != nil {
		log.Println("Failed to create gift card")
		return nil, err
	}

	var giftCard GiftCard
	internal.UnmarshalResponse(resp, &giftCard)

	return &giftCard, nil
}

func (s *Session) DisableGiftCard(giftCardId string) (*GiftCard, error) {
	endpoint := fmt.Sprintf("%s/%s", GiftCardsEndpoint, giftCardId)
	resp, err := internal.DeleteRequest(s.Secret, endpoint, nil)
	if err != nil {
		log.Println("Failed to disable giftcard")
		return nil, err
	}

	var giftCard GiftCard
	err = internal.UnmarshalResponse(resp, &giftCard)
	if err != nil {
		log.Println("Failed to unmarshal response")
		return nil, err
	}

	return &giftCard, nil
}

func (s *Session) TopUpGiftCard(giftCardId string) (*GiftCard, error) {
	endpoint := fmt.Sprintf("%s/%s", GiftCardsEndpoint, giftCardId)

	resp, err := internal.PutRequest(s.Secret, endpoint, nil)
	if err != nil {
		log.Println("Failed to create gift card")
		return nil, err
	}

	var giftCard GiftCard
	internal.UnmarshalResponse(resp, &giftCard)

	return &giftCard, nil
}
