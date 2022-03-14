package tebexgo

import (
	"github.com/itschip/tebexgo/internal"
	"log"
)

func (s *Session) GetAllBans() (*Bans, error) {
	res, err := internal.GetRequest(s.Secret, GetAllBansEndpoint)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	
	var bans Bans
	internal.UnmarshalResponse(res, &bans)
	
	return &bans, err
}

func (s *Session) CreateBan() {

}
