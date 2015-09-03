package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Verification struct {
	State       string `json:"state"`
	Description string `json:"description"`
	Context     string `json:"context"`
}

func CallVerification(url string, message Verification) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody(message)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func jsonBody(cb Verification) []byte {
	b, err := json.Marshal(cb)
	if err != nil {
		fmt.Println("error:", err)
	}
	return b
}
