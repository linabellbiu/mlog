package mlog

import "log"

type minLog struct {
	*logLv

	path string
}

type logLv struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}
