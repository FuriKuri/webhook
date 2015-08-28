package webhook

import (
	"encoding/json"
	"fmt"
)

type Verification struct {
	State       string `json:"state"`
	Description string `json:"description"`
	Context     string `json:"context"`
}

func CallVerification(url string, message Verification) {
	fmt.Println("Dummy call: " + string(response(message))[:])
}

func response(cb Verification) []byte {
	b, err := json.Marshal(cb)
	if err != nil {
		fmt.Println("error:", err)
	}
	return b
}
