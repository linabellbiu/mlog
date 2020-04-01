package mlog

import (
	"io"
	"log"
	"os"
)

const infoFileName = "info.log"
const warningFileName = "warning.log"
const errorFileName = "error.log"

var m = &minLog{
	logLv: nil,
}

func New(filePath string) *minLog {
	m.path = filePath
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
