package tebexgo

import (
	"fmt"
	"github.com/itschip/tebexgo/internal"
)

func (s *Session) GetAllCoupons() (*Coupons, error){
	resp, err := internal.GetRequest(s.Secret, CouponsEndpoint)
	if err != nil {
		return nil, err
	}
	
	var coupons Coupons
	err = internal.UnmarshalResponse(resp, &coupons)
	if err != nil {
		return nil, err
	}
	
	return &coupons, nil
}

func (s *Session) GetCoupon(couponId string) (*Coupon, error){
	endpoint := fmt.Sprintf("%s/%s", CouponsEndpoint, couponId)
	resp, err := internal.GetRequest(s.Secret, endpoint)
	if err != nil {
		return nil, err
	}
	
	var coupon Coupon
	err = internal.UnmarshalResponse(resp, &coupon)
	if err != nil {
		return nil, err
	}
	
	return &coupon, nil
}