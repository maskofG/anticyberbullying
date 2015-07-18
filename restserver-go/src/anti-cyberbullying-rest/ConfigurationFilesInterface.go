package main;


type ConfigurationFilesInterface interface{
	init(RestConfFile string, ESConfFile string, LoggerConfFile string)
    GetRestConfFile() (confFile string)
    SetRestConfFile(confFile string)
    GetESConfFile() (confFile string)
    SetESConfFile(confFile string)
    GetLoggerConfFile() (confFile string)
    SetLoggerConfFile(confFile string)
}