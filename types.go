package mlog

import "sync"

type MinLog struct {
	*logLv
	save int
	path string
}

type logLv struct {
	Info    *config
	Warning *config
	Error   *config
}

type config struct {
	On bool
	L  *sync.RWMutex
}
