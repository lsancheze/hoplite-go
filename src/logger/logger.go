//SPDX-FileCopyrightText: © 2023 3nets, Inc. <it@3nets.io>

package logger

import (
	"fmt"
	"time"
)

const (
	EMERGENCY int = iota
	ALERT
	CRITICAL
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

var logLevelname = []string{"emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"}
var logSource string
var logLevel int

func Init(source string) {
	logSource = source
	logLevel = INFO
}

func Level(level int) {
	logLevel = level
}

func Write(context string, level int, text string, json ...[]byte) {
	if level <= logLevel {
		timestamp := time.Now().Format(time.RFC3339)

		_json := ""
		if len(json) > 0 {
			_json = string(json[0])
		}
		n := len(_json)
		if n > 2 {
			fmt.Printf("{"+
				"\"timestamp\":\"%s\","+
				"\"source\":\"%s\","+
				"\"context\":\"%s\","+
				"\"level\":\"%s\","+
				"\"text\":\"%s\","+
				"%s"+
				"}\n",
				timestamp, logSource, context, logLevelname[level], text, _json[1:n-1])
		} else {
			fmt.Printf("{"+
				"\"timestamp\":\"%s\","+
				"\"source\":\"%s\","+
				"\"context\":\"%s\","+
				"\"level\":\"%s\","+
				"\"text\":\"%s\""+
				"}\n",
				timestamp, logSource, context, logLevelname[level], text)
		}
	}
}

func Print(context string, level int, format string, args ...interface{}) {
	if level <= logLevel {
		timestamp := time.Now().Format(time.RFC3339)

		fmt.Printf("{"+
			"\"timestamp\":\"%s\","+
			"\"source\":\"%s\","+
			"\"context\":\"%s\","+
			"\"level\":\"%s\","+
			"\"text\":\"%s\""+
			"}\n",
			timestamp, logSource, context, logLevelname[level], fmt.Sprintf(format, args...))
	}
}
