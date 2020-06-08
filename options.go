package mlog

import "os"

func (m *minLog) Info(msg string) {
	m.logLv.Info.Output(2, msg)
}

func (m *minLog) Warning(sign string, err error, msg string) {
	var error string
	if err != nil {
		error = err.Error()
	} else {
		error = ""
	}
	m.logLv.Warning.Output(2, "\n[Type]:"+sign+"\n[Error]:"+error+"\n[Msg]:"+msg)
}

func (m *minLog) Error(sign string, err error) {
	var error string
	if err != nil {
		error = err.Error()
	} else {
		error = ""
	}
	m.logLv.Error.Output(2, "\n[Type]:"+sign+"\n[Error]:"+error)
	os.Exit(1)
}
