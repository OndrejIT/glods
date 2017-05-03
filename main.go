package main

import (
	"github.com/ondrejit/glods/config"
	"github.com/ondrejit/glods/disk"
)

func init() {
	config.Setup()

}

func main() {
	disk.Check()
}
