package webhook

import (
	"log"
	"net/http"
)

type Repository struct {
	CommentCount    int64   `json:"comment_count"`
	DateCreated     float64 `json:"date_created"`
	Description     string  `json:"description"`
	Dockerfile      string  `json:"dockerfile"`
	FullDescription string  `json:"full_description"`
	IsOfficial      bool    `json:"is_official"`
	IsPrivate       bool    `json:"is_private"`
	IsTrusted       bool    `json:"is_trusted"`
	Name            string  `json:"name"`
	Namspace        string  `json:"namespace"`
	Owner           string  `json:"owner"`
	RepoName        string  `json:"repo_name"`
	RepoUrl         string  `json:"repo_url"`
	StarCount       uint32  `json:"star_count"`
	Status          string  `json:"status"`
}

type PushData struct {
	Images   []string `json:"images"`
	PushedAt float64  `json:"pushed_at"`
	Pusher   string   `json:"pusher"`
}

type Webhook struct {
	CallbackUrl string     `json:"callback_url"`
	PushData    PushData   `json:"push_data"`
	Repository  Repository `json:"repository"`
}

func Server(verifyFn func(Webhook) Verification) {
	http.HandleFunc("/webhook", WebhookHandler(verifyFn))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
