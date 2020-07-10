package mlog

import (
	"github.com/wangxudong123/assist"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"
)

func (m *minLog) Info(msg string) {

	if m.logLv.Info.On {
		filename := m.path + infoFileName
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic("打开" + filename + "日志文件失败:" + err.Error())
		}
		defer file.Close()
		log.New(io.MultiWriter(os.Stdout, file), "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile).Output(2, msg)
	}
}

func (m *minLog) Warning(sign string, err error, msg string) {
	var e string
	if err != nil {
		e = err.Error()
	} else {
		e = ""
	}
	if m.logLv.Warning.On {
		filename := m.path + warningFileName
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic("打开" + filename + "日志文件失败:" + err.Error())
		}
		defer file.Close()
		log.New(io.MultiWriter(os.Stdout, file), "[WARNING]: ", log.Ldate|log.Ltime|log.Llongfile).
			Output(2, "\n  Type:"+sign+"\n	Error:"+e+"\n	Msg:"+msg)
	}
}

func (m *minLog) Error(sign string, err error) {
	var e string
	if err != nil {
		e = err.Error()
	} else {
		e = ""
	}
	if m.logLv.Error.On {
		filename := m.path + errorFileName
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

func (m *minLog) delOutTime(path string) {
	l := len(path)
	if string(path[l-1]) != "/" {
		path = path + "/"
	}
	t := time.NewTicker(2 * time.Second)
	nowT := time.Now().Unix()
	for {
		if m.save > 0 {
			if files, err := ioutil.ReadDir(path); err == nil {
				for _, f := range files {
					reg := regexp.MustCompile(`^[0-9]+$`)
					result := reg.FindAllString(f.Name(), -1)
					if len(result) == 0 {
						continue
					}
					fileNameT := assist.StringToInt64(f.Name())
					nowT := assist.StringToInt64(assist.UnixToTimeFormats(nowT, "20060102"))
					if nowT-fileNameT > int64(m.save) {
						_ = os.RemoveAll(path + f.Name())
					}
				}
			}
		}
		<-t.C
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

func (m *minLog) Save(save int) {
	m.save = save
}
