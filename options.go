package mlog

func (m *minLog) Info(msg string) {
	m.logLv.Info.Output(2, msg)
}

func (m *minLog) Warning(msg string) {
	m.logLv.Warning.Output(2, msg)
}

func (m *minLog) Error(msg string) {
	m.logLv.Error.Output(2, msg)
}
