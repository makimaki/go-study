package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/sakamaki-albert/go-study/reply"
	"github.com/sakamaki-albert/go-study/util"
	"github.com/sakamaki-albert/go-study/webhook"
)

type LineConfig struct {
	ChannelAccessToken string `envconfig:"CHANNEL_ACCESS_TOKEN"`
	ChannelSecret      string `envconfig:"CHANNEL_SECRET"`
}

const (
	SIGNATURE_HEADER_NAME = "X-Line-Signature"
	REPLY_API_ENDPOINT    = "https://api.line.me/v2/bot/message/reply"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/webhook/:clientId/:integrationId", handleWebhook)

	r.Run()

	// request := `{
	// 	"events": [
	// 		{
	// 			"type": "message",
	// 			"timestamp": 1589347982,
	// 			"source": {
	// 				"type": "user",
	// 				"userId": "Uaaaa1"
	// 			},
	// 			"replyToken": "pleasereply",
	// 			"message": {
	// 				"type": "text",
	// 				"id": "text0001",
	// 				"text": "hoaaaaa"
	// 			}
	// 		},
	// 		{
	// 			"type": "message",
	// 			"timestamp": 1589347982,
	// 			"source": {
	// 				"type": "user",
	// 				"userId": "Uaaaa2"
	// 			},
	// 			"replyToken": "pleasereply",
	// 			"message": {
	// 				"type": "location",
	// 				"id": "loc0001",
	// 				"title": "大日本帝国",
	// 				"address": "日本列島",
	// 				"latitude": 35.362810,
	// 				"longitude": 138.731006
	// 			}
	// 		},
	// 		{
	// 			"type": "message",
	// 			"timestamp": 1589347982,
	// 			"source": {
	// 				"type": "user",
	// 				"userId": "Uaaaa3"
	// 			},
	// 			"replyToken": "pleasereply",
	// 			"message": {
	// 				"type": "location",
	// 				"id": "loc0002",
	// 				"latitude": 19.207428,
	// 				"longitude": -155.566406
	// 			}
	// 		},
	// 		{
	// 			"type": "postback",
	// 			"timestamp": 1589347982,
	// 			"source": {
	// 				"type": "user",
	// 				"userId": "Uaaaa4"
	// 			},
	// 			"replyToken": "pleasereply",
	// 			"postback": {
	// 				"data": "test"
	// 			}
	// 		},
	// 		{
	// 			"type": "follow",
	// 			"timestamp": 1589347982,
	// 			"source": {
	// 				"type": "user",
	// 				"userId": "Uaaaa5"
	// 			}
	// 		}
	// 	]
	// }`
}

func handleWebhook(c *gin.Context) {
	var lineConfig LineConfig
	if err := envconfig.Process("line", &lineConfig); err != nil {
		log.Fatalf("cannot read config from environment variables, %+v", err)
	}

	log.Println(lineConfig)

	buf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(buf)
	body := string(buf[0:n])

	log.Println(body)

	if !util.Validate(
		lineConfig.ChannelSecret,
		c.GetHeader(SIGNATURE_HEADER_NAME),
		body,
	) {
		log.Fatal(errors.New("aaaa"))
	}

	var req webhook.Request
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		log.Fatal(err)
	}
	// if err := c.BindJSON(&req); err != nil {
	// 	log.Fatal(err)
	// }

	for _, event := range req.Events {
		var res *reply.Request
		if err := handleEvent(event, &res); err != nil {
			log.Fatal(err)
		}
		if res != nil {
			sendReply(lineConfig, *res)
		}
	}
}

func handleEvent(e interface{}, res **reply.Request) error {
	switch e.(type) {
	case webhook.MessageEvent:
		var replyMessages []interface{}

		event := e.(webhook.MessageEvent)
		switch event.Message.(type) {
		case webhook.TextMessage:
			m := event.Message.(webhook.TextMessage)
			var replyText string
			if m.Text == "あなたは誰ですか？" {
				replyText = "私は golang で実装された何かです。"
			} else {
				replyText = fmt.Sprintf("%s ですね。わかります。", m.Text)
			}
			replyMessages = []interface{}{
				reply.TextMessage{
					Text: replyText,
				},
			}
		case webhook.LocationMessage:
			m := event.Message.(webhook.LocationMessage)
			targetTitle := "(不明)"
			if m.Title != nil {
				targetTitle = *m.Title
			}

			text := fmt.Sprintf("%s ですね。Google Maps で開きたいですか？", targetTitle)
			googleMapsUrl := fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%f,%f", m.Latitude, m.Longitude)
			replyMessages = []interface{}{
				reply.TemplateMessage{
					AltText: "これはだいたいてきすと",
					Template: reply.ButtonsTemplate{
						Text: text,
						Actions: []interface{}{
							reply.PostbackAction{
								Label:       "是非",
								DisplayText: "お願いします！",
								Data:        googleMapsUrl,
							},
						},
					},
				},
			}
		default:
			replyMessages = []interface{}{
				reply.TextMessage{
					Text: fmt.Sprintf("すみません。よくわかりませんのでダンプします。\n%+v", event.Message),
				},
			}
		}

		*res = &reply.Request{
			ReplyToken: event.ReplyToken,
			Messages:   replyMessages,
		}

	case webhook.PostbackEvent:
		event := e.(webhook.PostbackEvent)
		*res = &reply.Request{
			ReplyToken: event.ReplyToken,
			Messages: []interface{}{
				reply.TemplateMessage{
					AltText: "これはだいたいてきすと",
					Template: reply.ButtonsTemplate{
						Text: event.Postback.Data,
						Actions: []interface{}{
							reply.UriAction{
								Uri:   event.Postback.Data,
								Label: "ぐぐるまぷ",
							},
						},
					},
				},
			},
		}
	}

	return nil
}

func sendReply(config LineConfig, replyRequest reply.Request) error {
	buf, err := json.Marshal(replyRequest)

	if err != nil {
		return err
	}

	log.Printf("json:\n%s\n\n", buf)

	req, err := http.NewRequest(
		http.MethodPost,
		REPLY_API_ENDPOINT,
		bytes.NewBuffer(buf),
	)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.ChannelAccessToken))

	log.Println(config.ChannelAccessToken)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	log.Printf("res: %+v", response)

	return nil
}
