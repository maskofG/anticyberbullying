package main;


type ConfigurationFilesImpl struct {
	RestConfFile string
	ESConfFile   string
	LoggerConfFile string
}


func (conf *ConfigurationFilesImpl) init(RestConfFile string, ESConfFile string, LoggerConfFile string) {
    conf.RestConfFile = RestConfFile;
    conf.ESConfFile = ESConfFile;
    conf.LoggerConfFile = LoggerConfFile;
} 


func (conf *ConfigurationFilesImpl) GetRestConfFile() (confFile string) {
	confFile = conf.RestConfFile;
	return confFile;
}

func (conf *ConfigurationFilesImpl) SetRestConfFile(confFile string){
	conf.RestConfFile = confFile;
}


func (conf *ConfigurationFilesImpl) GetESConfFile() (confFile string) {
	confFile = conf.ESConfFile;
	return confFile;
}

func (conf *ConfigurationFilesImpl) SetESConfFile(confFile string){
	conf.ESConfFile = confFile;
}



func (conf *ConfigurationFilesImpl) GetLoggerConfFile() (confFile string) {
	confFile = conf.LoggerConfFile;
	return confFile;
}


func (conf *ConfigurationFilesImpl) SetLoggerConfFile(confFile string){
	conf.LoggerConfFile = confFile;
}