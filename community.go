package tebexgo

import (
	"fmt"

	"github.com/itschip/tebexgo/internal"
)

func (s *Session) GetAllCommunityGoals() (*[]CommunityGoal, error) {
	resp, err := internal.GetRequest(s.Secret, CommunityEndpoint)
	if err != nil {
		return nil, err
	}

	var communityGoals []CommunityGoal
	err = internal.UnmarshalResponse(resp, &communityGoals)
	if err != nil {
		return nil, err
	}

	return &communityGoals, nil
}

func (s *Session) RetrieveCommunityGoal(communityGoal int) (*CommunityGoal, error) {
	endpoint := fmt.Sprintf("%s/%s", CommunityEndpoint, communityGoal)
	resp, err := internal.GetRequest(s.Secret, endpoint)
	if err != nil {
		return nil, err
	}

	var communityGoals CommunityGoal
	err = internal.UnmarshalResponse(resp, &communityGoals)
	if err != nil {
		return nil, err
	}

	return &communityGoals, nil
}
