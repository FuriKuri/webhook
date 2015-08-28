package webhook

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WebhookHandler(verifyFn func(Webhook) Verification) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var webhook Webhook
		err := decoder.Decode(&webhook)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "{\"state\":\"ok\"}")
		go CallVerification(webhook.CallbackUrl, verifyFn(webhook))
	}
}
