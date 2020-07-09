package mlog

import (
	"os"
	"runtime"
	"syscall"
	"time"
)

func getCreateTime(path string) int64 {
	osType := runtime.GOOS
	fileInfo, _ := os.Stat(path)
	if osType == "linux" {
		//Linux下文件创建时间
		// Sys()返回的是interface{}，不同平台需要的类型不一样，linux上为*syscall.Stat_t
		statT := fileInfo.Sys().(*syscall.Stat_t)
		return statT.Ctim.Sec
	}
	return time.Now().Unix()
}
