package mlog

import (
	"time"
)

const infoFileName = "info.log"
const errorFileName = "error.log"
const warningFileName = "warning.log"

var m = &minLog{
	logLv: nil,
}

type saveTime int32

const (
	SaveDay   saveTime = 0
	SaveWeek  saveTime = 1
	SaveMonth saveTime = 2
	SaveYear  saveTime = 3
)

const (
	//LOGPATH  LOGPATH/time.Now().Format(FORMAT)/*.log
	LOGPATH = "/"
	//FORMAT .
	FORMAT = "20060102"
)

func New(filePath string) *minLog {
	m.path = filePath
	m.logLv = &logLv{
		Info:    &config{On: true},
		Warning: &config{On: true},
		Error:   &config{On: true},
	}
	return m
}

func (m *minLog) parse() string {
	path := m.path + LOGPATH + time.Now().Format(FORMAT) + "/"
	l := len(path)
	if string(path[l-1]) != "/" {
		path = path + "/"
	}
	return path
}
