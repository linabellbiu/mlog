package mlog

import (
	"io"
	"log"
	"os"
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

var path = LOGPATH + time.Now().Format(FORMAT) + "/"

func New(filePath string) *minLog {
	//m.save = save
	m.parse(filePath + path)
	if err := cDir(m.path); err != nil {
		panic("创建目录" + m.path + "失败:" + err.Error())
	}
	m.logLv = m.newLogLv()
	return m
}

func (m *minLog) newLogLv() *logLv {
	return &logLv{
		Info:    log.New(io.MultiWriter(os.Stderr, cFile(m.path+infoFileName)), "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile),
		Warning: log.New(io.MultiWriter(os.Stderr, cFile(m.path+warningFileName)), "[WARNING]: ", log.Ldate|log.Ltime|log.Llongfile),
		Error:   log.New(io.MultiWriter(os.Stderr, cFile(m.path+errorFileName)), "[ERROR]: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile),
	}
}

func (m *minLog) parse(path string) {
	m.path = path
	l := len(path)
	if string(path[l-1]) != "/" {
		m.path = path + "/"
	}
}
