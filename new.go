package mlog

import (
	"os"
	"time"
)

const infoFileName = "info.log"
const errorFileName = "error.log"
const warningFileName = "warning.log"

var m = &minLog{
	logLv: nil,
}

const (
	//LOGPATH  LOGPATH/time.Now().Format(FORMAT)/*.log
	LOGPATH = "/"
	//FORMAT .
	FORMAT = "20060102"
)

func New(filePath string) *minLog {
	if filePath == "" {
		filePath = "./"
	}

	m.logLv = &logLv{
		Info:    &config{On: true},
		Warning: &config{On: true},
		Error:   &config{On: true},
	}

	if err := m.parse(filePath); err != nil {
		panic("创建目录" + m.path + "失败:" + err.Error())
	}

	go m.delOutTime(filePath)

	return m
}

func (m *minLog) parse(path string) error {
	m.path = path + LOGPATH + time.Now().Format(FORMAT) + "/"
	l := len(m.path)
	if string(m.path[l-1]) != "/" {
		m.path = m.path + "/"
	}
	return os.MkdirAll(m.path, os.ModePerm)
}
