package mlog

import (
	"os"
	"time"
)

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
