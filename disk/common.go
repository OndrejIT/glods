package disk

import (
	"fmt"
	"os"
	conf "github.com/spf13/viper"
	log "github.com/Sirupsen/logrus"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func MsgFormater(path string, status DiskStatus) string {
	var hostname string

	if h := conf.GetString("hostname"); h == "" {
		hostname, _ = os.Hostname()
	} else {
		hostname = h
	}

	msg := fmt.Sprintf(
		"--- WARNING: Low disk space ---\n"+
			"Hostname: %s\n"+
			"Path: %s\n"+
			"Size: %.2f GB\n"+
			"Used: %.2f GB (%.2f %%)\n"+
			"Free: %.2f GB (%.2f %%)",
		hostname,
		path,
		status.Size/GB,
		status.Used/GB,
		status.PercentUsed,
		status.Free/GB,
		status.PercentFree,
	)

	if conf.GetBool("silent") != true {
		log.Info(msg)
	}

	return msg
}
