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

func (m *MinLog) Info(msg string) {
	m.logLv.Info.L.RLock()
	ok := m.logLv.Info.On
	m.logLv.Info.L.RUnlock()
	if ok {
		filename, err := m.parse()
		if err != nil {
			panic("创建" + filename + "日志文件失败:" + err.Error())
		}
		filename = filename + infoFileName
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic("打开" + filename + "日志文件失败:" + err.Error())
		}
		defer file.Close()
		log.New(io.MultiWriter(os.Stdout, file), "[INFO]: ", log.Ldate|log.Lmicroseconds|log.Lshortfile).Output(2, msg)
	}
}

func (m *MinLog) Warning(sign string, err error, msg string) {
	var e string
	if err != nil {
		e = err.Error()
	} else {
		e = ""
	}
	m.logLv.Warning.L.RLock()
	ok := m.logLv.Warning.On
	m.logLv.Warning.L.RUnlock()
	if ok {
		filename, err := m.parse()
		if err != nil {
			panic("创建" + filename + "日志文件失败:" + err.Error())
		}
		filename = filename + warningFileName
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic("打开" + filename + "日志文件失败:" + err.Error())
		}
		defer file.Close()
		log.New(io.MultiWriter(os.Stdout, file), "[WARNING]: ", log.Ldate|log.Ltime|log.Llongfile).
			Output(2, "\n  Type:"+sign+"\n	Error:"+e+"\n	Msg:"+msg)
	}
}

func (m *MinLog) Error(sign string, err error) {
	var e string
	if err != nil {
		e = err.Error()
	} else {
		e = ""
	}
	m.logLv.Error.L.RLock()
	ok := m.logLv.Error.On
	m.logLv.Error.L.RUnlock()
	if ok {
		filename, err := m.parse()
		if err != nil {
			panic("创建" + filename + "日志文件失败:" + err.Error())
		}
		filename = filename + errorFileName
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

func (m *MinLog) delOutTime(path string) {
	l := len(path)
	if string(path[l-1]) != "/" {
		path = path + "/"
	}
	t := time.NewTicker(2 * time.Second)
	nowT := time.Now().Unix()
	for {
		m.l.RLock()
		save := m.save
		m.l.RUnlock()
		if save >= 1 {
			if files, err := ioutil.ReadDir(path); err == nil {
				for _, f := range files {
					reg := regexp.MustCompile(`^[0-9]+$`)
					result := reg.FindAllString(f.Name(), -1)
					if len(result) == 0 {
						continue
					}
					fileNameT := assist.StringToInt64(f.Name())
					nowT := assist.StringToInt64(assist.UnixToTimeFormats(nowT, "20060102"))
					if nowT-fileNameT > int64(save-1) {
						_ = os.RemoveAll(path + f.Name())
					}
				}
			}
		}
		<-t.C
	}
}

func (m *MinLog) SetInfoOutPut(status bool) {
	m.logLv.Info.L.Lock()
	m.logLv.Info.On = status
	m.logLv.Info.L.Unlock()
}

func (m *MinLog) SetWarningOutPut(status bool) {
	m.logLv.Warning.L.Lock()
	m.logLv.Warning.On = status
	m.logLv.Warning.L.Unlock()
}

func (m *MinLog) SetErrorOutPut(status bool) {
	m.logLv.Error.L.Lock()
	m.logLv.Error.On = status
	m.logLv.Error.L.Unlock()
}

//保存几天
func (m *MinLog) Save(save int) {
	m.l.Lock()
	m.save = save
	m.l.Unlock()
}
