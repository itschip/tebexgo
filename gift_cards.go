package tebexgo

import (
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
