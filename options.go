package mlog

import (
	"io"
	"log"
	"os"
)

func (m *minLog) Info(msg string) {
	if path, err := cDir(); err != nil {
		panic("创建目录" + path + "失败:" + err.Error())
	} else {
		if m.logLv.Info.On {
			filename := path + infoFileName
			file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				panic("打开" + filename + "日志文件失败:" + err.Error())
			}
			defer file.Close()

			log.New(io.MultiWriter(os.Stdout, file), "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile).Output(2, msg)
		}
	}
}

func (m *minLog) Warning(sign string, err error, msg string) {
	if path, errs := cDir(); errs != nil {
		panic("创建目录" + path + "失败:" + errs.Error())
	} else {
		var e string
		if err != nil {
			e = err.Error()
		} else {
			e = ""
		}
		if m.logLv.Warning.On {
			filename := path + warningFileName
			file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				panic("打开" + filename + "日志文件失败:" + err.Error())
			}
			defer file.Close()
			log.New(io.MultiWriter(os.Stdout, file), "[WARNING]: ", log.Ldate|log.Ltime|log.Llongfile).
				Output(2, "\n  Type:"+sign+"\n	Error:"+e+"\n	Msg:"+msg)
		}
	}
}

func (m *minLog) Error(sign string, err error) {
	if path, errs := cDir(); errs != nil {
		panic("创建目录" + path + "失败:" + err.Error())
	} else {
		var e string
		if err != nil {
			e = err.Error()
		} else {
			e = ""
		}
		if m.logLv.Error.On {
			filename := path + errorFileName
			file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				panic("打开" + filename + "日志文件失败:" + err.Error())
			}
			defer file.Close()
			log.New(io.MultiWriter(os.Stdout, file), "[ERROR]: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile).
				Output(2, "\nType:"+sign+"\nError:"+e)
			panic(e)
		}
	}
}

func (m *minLog) SetInfoOutPut(status bool) {
	m.logLv.Info = &config{On: status}
}

func (m *minLog) SetWarningOutPut(status bool) {
	m.logLv.Warning = &config{On: status}
}

func (m *minLog) SetErrorOutPut(status bool) {
	m.logLv.Error = &config{On: status}
}
