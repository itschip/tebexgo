package tebexgo

import (
	"fmt"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	s := New("secret")
	pkg, err := s.GetPackage("pkg_id")
	if err != nil {
		log.Println("Failed to get package")
	}
	
	fmt.Println(pkg.Name)
}