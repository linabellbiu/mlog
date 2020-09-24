package mlog

import (
	"os"
	"runtime"
	"syscall"
	"testing"
	"time"
)

type Win32FileAttributeData struct {
	FileAttributes uint32
	CreationTime   Filetime
	LastAccessTime Filetime
	LastWriteTime  Filetime
	FileSizeHigh   uint32
	FileSizeLow    uint32
}
type Filetime struct {
	LowDateTime  uint32
	HighDateTime uint32
}

func TestgetCreateTime(path string, t *testing.T) int64 {
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

func Test_main(t *testing.T) {

}
