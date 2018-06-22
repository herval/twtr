package util

import (
	"log"
	"os"
)

var (
	Log *log.Logger
)

func InitLogger() {
	// set location of log file
	var logpath = "info.log"

	var file, err1 = os.Create(logpath)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
