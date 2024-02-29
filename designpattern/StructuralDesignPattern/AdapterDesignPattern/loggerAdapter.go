package adapterdesignpattern

type LoggerAdapter struct {
	Adapter *LegacyLogger
}

func (l *LoggerAdapter) Log(message string) {
	l.Adapter.LogMessage(message)
}
