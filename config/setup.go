package config

import (
	log "github.com/sirupsen/logrus"
	conf "github.com/spf13/viper"
)

func Setup() {
	FlagParser()

	if conf.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
		text_formatter := new(log.TextFormatter)
		text_formatter.TimestampFormat = "2006-01-02 15:04:05"
		log.SetFormatter(text_formatter)
		text_formatter.FullTimestamp = true
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}

	conf.SetConfigName("glods")
	conf.SetEnvPrefix("glods")
	conf.AutomaticEnv()

	conf.ReadInConfig()

}
