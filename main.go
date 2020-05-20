package main

import (
	// "net/http"
	"encoding/json"

	// "github.com/gin-gonic/gin"

	"log"

	"github.com/sakamaki-albert/go-study/reply"
	"github.com/sakamaki-albert/go-study/webhook"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.Run()

	request := `{
		"events": [
			{
				"type": "message",
				"timestamp": 1589347982,
				"source": {
					"type": "user",
					"userId": "Uaaaa1"
				},
				"replyToken": "pleasereply",
				"message": {
					"type": "text",
					"id": "text0001",
					"text": "hoaaaaa"
				}
			},
			{
				"type": "message",
				"timestamp": 1589347982,
				"source": {
					"type": "user",
					"userId": "Uaaaa2"
				},
				"replyToken": "pleasereply",
				"message": {
					"type": "location",
					"id": "loc0001",
					"title": "大日本帝国",
					"address": "日本列島",
					"latitude": 35.362810,
					"longitude": 138.731006
				}
			},
			{
				"type": "message",
				"timestamp": 1589347982,
				"source": {
					"type": "user",
					"userId": "Uaaaa3"
				},
				"replyToken": "pleasereply",
				"message": {
					"type": "location",
					"id": "loc0002",
					"latitude": 19.207428,
					"longitude": -155.566406
				}
			},
			{
				"type": "postback",
				"timestamp": 1589347982,
				"source": {
					"type": "user",
					"userId": "Uaaaa4"
				},
				"replyToken": "pleasereply",
				"postback": {
					"data": "test"
				}
			},
			{
				"type": "follow",
				"timestamp": 1589347982,
				"source": {
					"type": "user",
					"userId": "Uaaaa5"
				}
			}
		]
	}`

	var req webhook.Request
	if err := json.Unmarshal([]byte(request), &req); err != nil {
		log.Fatal(err)
	}

	for _, event := range req.Events {
		switch event.(type) {
		case webhook.MessageEvent:
			message := event.(webhook.MessageEvent).Message
			switch message.(type) {
			case webhook.TextMessage:
				m := message.(webhook.TextMessage)
				log.Printf("%s, テキストです\n", m.Text)
			case webhook.LocationMessage:
				m := message.(webhook.LocationMessage)
				log.Printf("%f, %f, 位置情報です\n", m.Latitude, m.Longitude)
			}

		case webhook.PostbackEvent:
			log.Println("ポストバック")
		}
	}

	res := reply.Request{
		ReplyToken: "a",
		Messages: []interface{}{
			reply.TextMessage{
				Text: "",
			},
		},
	}

	buf, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("json:\n%s\n\n", buf)
}
