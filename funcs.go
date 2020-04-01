package mlog

import "os"

func cFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("打开" + filename + "日志文件失败:" + err.Error())
	}
	return file
}
