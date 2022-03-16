package tebexgo

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := New("be25ccea31449768e8961afade6cfc166a46e0ee")

	data, _ := s.CreateBan(&BanInput{Reason: "fuck u bitch", Ip: "", User: "610495"})
	fmt.Println(data.Reason)
}
