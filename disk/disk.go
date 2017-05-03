package disk

import (
	"syscall"
	"strings"
	"github.com/ondrejit/glods/telegram"
	conf "github.com/spf13/viper"
	log "github.com/Sirupsen/logrus"
)

type DiskStatus struct {
	Size        float64 `json:"size"`
	Used        float64 `json:"used"`
	PercentUsed float64 `json:"percent_used"`
	Free        float64 `json:"free"`
	PercentFree float64 `json:"percent_free"`
}

func Status(path string) (DiskStatus, error) {
	fs := syscall.Statfs_t{}

	err := syscall.Statfs(path, &fs)
	if err != nil {
		return DiskStatus{}, err
	}

	size := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bavail * uint64(fs.Bsize)
	used := size - free

	stats := DiskStatus{}
	stats.Size = float64(size)
	stats.Used = float64(used)
	stats.Free = float64(free)
	stats.PercentFree = 100 - (stats.Size-stats.Free)/stats.Size*100
	stats.PercentUsed = (stats.Size - stats.Free) / stats.Size * 100

	return stats, nil
}

func Check() {
	pathConf := conf.GetString("path")
	path := strings.Split(pathConf, ",")

	for p := range path {
		status, err := Status(path[p])
		if err != nil {
			log.Fatalf("[Disk] %s", err)
		}

		warning := conf.GetInt("warning")
		if status.Free/GB < float64(warning) {
			telegram.Send(MsgFormater(path[p], status))
		}
	}
}
