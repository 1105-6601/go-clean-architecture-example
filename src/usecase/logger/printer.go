package logger

type Printer interface {
	Error    (text string)
	ErrorF   (format string, a ...interface{})
	Success  (text string)
	SuccessF (format string, a ...interface{})
	Info     (text string)
	InfoF    (format string, a ...interface{})
	Warn     (text string)
	WarnF    (format string, a ...interface{})
	Other    (text string)
	OtherF   (format string, a ...interface{})
}
