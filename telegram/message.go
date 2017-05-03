package telegram

import (
	"gopkg.in/telegram-bot-api.v4"
	log "github.com/Sirupsen/logrus"
	conf "github.com/spf13/viper"
)

func Send(msg string) {
	token := conf.GetString("token")
	if token == "" {
		return
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Errorf("[Telegram] %s", err)
		return
	}

	if conf.GetBool("verbose") {
		bot.Debug = true
	}

	chat := conf.GetInt64("chat")
	tapi := tgbotapi.NewMessage(chat, msg)

	bot.Send(tapi)
}
