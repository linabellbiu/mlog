package mlog

import (
	"os"
	"runtime"
	"syscall"
	"time"
)

func getCreateTime(path string) int64 {
	osType := runtime.GOOS
	fileInfo, err := os.Stat(path)
	if err != nil {
		return time.Now().Unix()
	}
	if osType == "darwin" {
		//Linux下文件创建时间
		// Sys()返回的是interface{}，不同平台需要的类型不一样，linux上为*syscall.Stat_t
		statT := fileInfo.Sys().(*syscall.Stat_t)
		return statT.Ctimespec.Sec
	}
	return time.Now().Unix()
}
