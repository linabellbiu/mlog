package mlog

func (m *minLog) Info(msg string) {
	m.logLv.Info.Println(msg)
}

func (m *minLog) Warning(msg string) {
	m.logLv.Warning.Println(msg)
}

func (m *minLog) Error(msg string) {
	m.logLv.Error.Println(msg)
}
