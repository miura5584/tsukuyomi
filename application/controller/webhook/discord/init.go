package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"tsukuyomi/controller/webhook/discord/v1"
	"tsukuyomi/controller/webhook/discord/v2"
	"tsukuyomi/logger"
)

var (
	Token   = "OTI0NTEwNTk1MDkxNDE1MDQw.Ycfnlg.TFM3Rl7za2QK-bLnCmsRES-3DQo"
	BotName = "Tsukuyomi"
	//stopBot           = make(chan bool)
	//vcsession         *discordgo.VoiceConnection
	helloWorld = "!helloworld"
	//ChannelVoiceJoin  = "!vcJoin"
	//ChannelVoiceLeave = "!vcleave"
)

func init() {
	discord, err := discordgo.New(Token)
	if err != nil {
		logger.Log.Error("failed initialize discord api")
	}

	discord.AddHandler(onMessageCreate)
	err = discord.Open()
	if err != nil {
		// TODO:
		fmt.Println(err)
	}

	defer func(discord *discordgo.Session) {
		err := discord.Close()
		if err != nil {
			log.Println("failed close discord session")
		}
	}(discord)

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot
	return
}

func AttachEndpoint(webhook *gin.RouterGroup) {
	discordEndpoint := webhook.Group("discord")
	{
		discordEndpoint.POST("/v1", v1.WebhookV1)
		discordEndpoint.POST("/v2", v2.WebhookV2)
	}
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	switch {
	case strings.HasPrefix(m.Content, fmt.Sprintf("%s %s", BotName, helloWorld)):
		sendMessage(s, m.ChannelID, "Hello World!!")
		//case strings.HasPrefix(m.Content, fmt.Sprintf("%s %s", BotName, ChannelVoiceJoin)):
		//	//今いるサーバーのチャンネル情報の一覧を喋らせる処理を書いておきますね
		//	c, err := s.State.Channel(m.ChannelID) //チャンネル取得
		//	if err != nil {
		//		log.Fatalf("error")
		//	}
		//	guildChannels, _ := s.GuildChannels(c.GuildID)
		//	var sendText string
		//	for _, a := range guildChannels {
		//		sendText += fmt.Sprintf("%vチャンネルの%v(IDは%v)\n", a.Type, a.Name, a.ID)
		//	}
		//	sendMessage(s, c, sendText) // チャンネルの名前、ID、タイプ(通話orテキスト)をBOTが話す
		//
		//	//VOICE CHANNEL IDには、botを参加させたい通話チャンネルのIDを代入してください
		//	//コメントアウトされた上記の処理を使うことでチャンネルIDを確認できます
		//	cvsession, _ = s.ChannelVoiceJoin(c)
		//case strings.HasPrefix(m.Content, fmt.Sprintf("%s %s", BotName, ChannelVoiceLeave)):
		//	vcsession.Disconnect()
	}
}

func onVoiceReceived(cv *discordgo.VoiceConnection, vs *discordgo.VoiceSpeakingUpdate) {
	log.Print("おれのおれのおれの話をきけ!")
}

func sendMessage(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)
	if err != nil {
		log.Print("failed sending message")
	}
}
