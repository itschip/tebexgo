package internal

import (
	"encoding/json"
	"log"
)

func UnmarshalResponse(response []byte, v interface{}) error {
	err := json.Unmarshal(response, &v)
	if err != nil {
		log.Println("Failed to marshal response")
		return err
	}
	
	return nil
}