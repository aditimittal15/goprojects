package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	db "goprojects/articleStore/server/handler"
	"runtime"
	"strings"
)

func initLog() {
	formatter := &log.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(formatter)
}
func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func init() {

	initLog()
	db.CreateDbConnection()
	db.CreateDBSchema()
}
