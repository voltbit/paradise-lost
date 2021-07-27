package main

/*
Implementation based on:
https://golangcode.com/send-slack-messages-without-a-library/
https://github.com/ashwanthkumar/slack-go-webhook/blob/master/main.go
*/

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type SlackRequestBody struct {
	Parse     string `json:"parse,omitempty"`
	Username  string `json:"username,omitempty"`
	IconUrl   string `json:"icon_url,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
	Channel   string `json:"channel,omitempty"`
	Text      string `json:"text,omitempty"`
	LinkNames string `json:"link_names,omitempty"`
	// Attachments []Attachment `json:"attachments,omitempty"`
	UnfurlLinks bool `json:"unfurl_links,omitempty"`
	UnfurlMedia bool `json:"unfurl_media,omitempty"`
	Markdown    bool `json:"mrkdwn,omitempty"`
}

func main() {
	webhookURL := os.Getenv("SLACK_WEBHOOK_DEV")
	req := NewRequestBody(MakeVaultPGPKeysMessageText())
	err := SendSlackNotification(webhookURL, req)
	if err != nil {
		log.Fatal(err)
	}
}

func MakeVaultPGPKeysMessageText() string {
	keys := []string{"zs6ZQEexLge6hP5qhWA8P7hqkWEKuQ4RgAGE4N9k", "Bo0k6XmE9ppDT8uHcUXepBLks88nMhScu2vEKbrq", "kiSVEfHNpQ67CJxk7bWsgyViHohErlzyzCD7ogeP", "uMOdFSr2HN5qhwR3lYGxcMSlE1pNA5xRw5a5k6Wu", "ZZXhTSGzpSCkUfn9tOxUGXhTSGzpSCkUfn9tOxUG"}
	formatedKeyText := ""
	for _, key := range keys {
		formatedKeyText += fmt.Sprintf("%s\n", key)
	}
	formatedKeyText = fmt.Sprintf("```%s```", formatedKeyText)
	return fmt.Sprintf("`user: %s`\n%s", "voltbit", formatedKeyText)
}

func NewRequestBody(message string) *SlackRequestBody {
	return &SlackRequestBody{
		Username: "VaultGoopher",
		Text:     message,
		IconUrl:  "https://jonathanmh.com/wp-content/uploads/2018/01/jonathan-gopher.png",
		Markdown: true,
	}
}

// SendSlackNotification will post to an 'Incoming Webook' url setup in Slack Apps. It accepts
// some text and the slack channel is saved within Slack.
func SendSlackNotification(webhookUrl string, body *SlackRequestBody) error {
	slackBody, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}
