package mlog

type minLog struct {
	*logLv
	save saveTime
	path string
}

type logLv struct {
	Info    *config
	Warning *config
	Error   *config
}

type config struct {
	On bool
}
