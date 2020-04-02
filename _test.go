package mlog_test

import (
	"mlog"
)

func test() {
	m := mlog.New("./log")
	m.Info("这是一个info log")
}
