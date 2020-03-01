package main

import (
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	db "goprojects/articleStore/server/handler"
)

func initLog() {
	formatter := &log.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
	}
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(formatter)
}

func init() {

	initLog()
	db.CreateDbConnection()
	db.CreateDBSchema()
}
