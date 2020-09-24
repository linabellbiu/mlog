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
	if osType == "windows" {
		//Win下文件创建时间
		fileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
		nanoseconds := fileSys.CreationTime.Nanoseconds() // 返回的是纳秒
		createTime := nanoseconds / 1e9                   //秒
		return createTime
	}
	return time.Now().Unix()
}
