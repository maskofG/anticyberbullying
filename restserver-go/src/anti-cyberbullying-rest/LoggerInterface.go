package main;


type LoggerInterface interface {
	init(ConfigurationFileListPtr ConfigurationFilesInterface, filename string)
	Info(message ...interface{})
	Warn(message ...interface{})
	Debug(message ...interface{})
	Error(message ...interface{})
	Critical(message ...interface{})
	Trace(message ...interface{})
}