# Golang low disk space warning [![Build Status](https://travis-ci.org/OndrejIT/glods.svg?branch=master)](https://travis-ci.org/OndrejIT/glods)
Work with [telegram](https://telegram.org/)!

### Install
  - or: go get

### Run
 - go build
 - glods --token telegram_bot_token --chat telegram_chat_id --path /srv,/mnt/data --warning 200

### Usage
  Usage of ./glods:
  * -i, --chat int          Set telegram chat ID.
  * -c, --config string     Set config path. (default ".")
  * -d, --debug             Enable debug mode.
  * -n, --hostname string   Set hostname.
  * -p, --path string       Set disk path. ("/root,/srv") (default "/")
  * -s, --silent            Set silent mode in terminal.
  * -t, --token string      Set telegram bot token.
  * -v, --verbose           Enable verbose mode.
  * -w, --warning int       Set warning limit. (in GB) (default 100)
