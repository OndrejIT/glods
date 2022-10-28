package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
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
