package mlog

import (
	"os"
	"time"
)

func cFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("打开" + filename + "日志文件失败:" + err.Error())
	}
	return file
}

func cDir() (string, error) {
	path := m.parse()
	return path, os.MkdirAll(path, os.ModePerm)
}

func cPackPage(path string) bool {
	now := time.Now().Unix()
	ct := getCreateTime(path)
	switch m.save {
	case SaveDay:
		if ct+3600*24 > now {
			return false
		}
	case SaveWeek:
		if ct+7*3600*24 > now {
			return false
		}
	case SaveMonth:
		if ct+30*3600*24 > now {
			return false
		}
	case SaveYear:
		if ct+365*3600*24 > now {
			return false
		}
	default:
		if ct+3600*24 > now {
			return false
		}
	}
	return true
}
