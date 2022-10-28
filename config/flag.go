package config

import (
	flag "github.com/spf13/pflag"
	conf "github.com/spf13/viper"
)

func FlagParser() {
	flag.StringP("path", "p", "/", "Set disk path. (\"/root,/srv\")")
	flag.StringP("hostname", "n", "", "Set hostname.")
	flag.IntP("warning", "w", 100, "Set warning limit. (in GB or percent)")
	flag.Bool("percent", false, "Set warning limit in percent mode.")
	flag.StringP("token", "t", "", "Set telegram bot token.")
	flag.Int64P("chat", "i", 0, "Set telegram chat ID.")
	flag.StringP("config", "c", ".", "Set config path.")
	flag.BoolP("silent", "s", false, "Set silent mode in terminal.")
	flag.BoolP("debug", "d", false, "Enable debug mode.")
	flag.BoolP("verbose", "v", false, "Enable verbose mode.")
	flag.Parse()

	flagToConfig()
}

func flagToConfig() {
	conf.BindPFlag("hostname", flag.Lookup("hostname"))
	conf.BindPFlag("warning", flag.Lookup("warning"))
	conf.BindPFlag("percent", flag.Lookup("percent"))
	conf.BindPFlag("token", flag.Lookup("token"))
	conf.BindPFlag("chat", flag.Lookup("chat"))
	conf.BindPFlag("path", flag.Lookup("path"))
	conf.BindPFlag("config", flag.Lookup("config"))
	conf.BindPFlag("silent", flag.Lookup("silent"))
	conf.BindPFlag("debug", flag.Lookup("debug"))
	conf.BindPFlag("verbose", flag.Lookup("verbose"))
}
