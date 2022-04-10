package v1

import (
	"context"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"tsukuyomi/ent"
)

func eventFollow(event *linebot.Event) {
	ctx := context.Background()
	switch event.Source.Type {
	case linebot.EventSourceTypeUser:
		_, err := ent.Database.LineUser.Create().SetID(event.Source.UserID).Save(ctx)
		if err != nil {

		}
	case linebot.EventSourceTypeGroup:
	case linebot.EventSourceTypeRoom:
	}
}

func eventUnFollow(event *linebot.Event) {
	ctx := context.Background()
	switch event.Source.Type {
	case linebot.EventSourceTypeUser:
		_, err := ent.Database.LineUser.UpdateOneID(event.Source.UserID).SetIsActive(false).Save(ctx)
		if err != nil {

		}
	case linebot.EventSourceTypeGroup:
	case linebot.EventSourceTypeRoom:
	}
}
