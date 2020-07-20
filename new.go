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

	m.path = filePath

	m.logLv = &logLv{
		Info:    &config{On: true},
		Warning: &config{On: true},
		Error:   &config{On: true},
	}

	//if err := m.parse(filePath); err != nil {
	//	panic("创建目录" + m.path + "失败:" + err.Error())
	//}

	go m.delOutTime(filePath)

	return m
}

func (m *minLog) parse() (string, error) {
	var path string
	l := len(m.path)
	if string(m.path[l-1]) != "/" {
		path = m.path + "/"
	}
	path = path + LOGPATH + time.Now().Format(FORMAT) + "/"

	return path, os.MkdirAll(path, os.ModePerm)
}
