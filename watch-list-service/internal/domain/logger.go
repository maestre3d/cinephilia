package domain

// LogField set custom fields to a log
type LogField struct {
	Key   string
	Value interface{}
}

// Logger prints given data to a source like stdout, file system, etc...
//	@Port
type Logger interface {
	Debug(msg string, args ...LogField)
	Info(msg string, args ...LogField)
	Warn(msg string, args ...LogField)
	Error(msg string, args ...LogField)
}
