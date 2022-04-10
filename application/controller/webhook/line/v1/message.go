package v1

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"tsukuyomi/controller/webhook/line"
)

func eventMessage(event *linebot.Event) {
	var err error
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		text := ""
		if _, err = line.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text)).Do(); err != nil {

		}
	case *linebot.StickerMessage:
		replyMessage := fmt.Sprintf(
			"sticker id is %s, stickerResourceType is %s",
			message.StickerID,
			message.StickerResourceType,
		)
		if _, err = line.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		}
	}
}
