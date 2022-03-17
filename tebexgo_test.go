package tebexgo

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := New("")

	data, _ := s.CreateBan(&BanInput{Reason: "some reason", Ip: "", User: "some uid"})
	fmt.Println(data.Reason)
}
