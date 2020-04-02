package mlog

import (
	"io"
	"log"
	"os"
)

const infoFileName = "info.log"
const errorFileName = "error.log"
const warningFileName = "warning.log"

var m = &minLog{
	logLv: nil,
}

func New(filePath string) *minLog {
	m.parse(filePath)
	if err := cDir(m.path); err != nil {
		panic("创建目录" + m.path + "失败")
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
