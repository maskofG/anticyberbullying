package main


import (
	"strings"
	"regexp"
	"strconv"
)

import log "github.com/cihub/seelog"
import "github.com/revel/config"


var LogFileSizePattern = "[1-9][0-9]*|[B]|[kK]B|[mM]B|[gG]B|[tT]B"

type LogRotationMetadata struct {
	RotationType  string
	MaxRollNumber int
	Properties map[string]string
}

type LoggerImpl struct {
	ConfigurationFileListPtr ConfigurationFilesInterface;
	IsLogRotationEnabled  bool
	LogsFolder            string
	LogFiles              []string
	LogMinLevel           string
	LogMaxLevel           string
	LogFormat             string
	LogConfigurationFile  string
	LogRotationAttributes LogRotationMetadata
	Logger                log.LoggerInterface
	RelativeFileName      string
}


func (logger *LoggerImpl) init(ConfigurationFileListPtr ConfigurationFilesInterface, filename string) {
	logger.RelativeFileName = filename;
	logger.ConfigurationFileListPtr = ConfigurationFileListPtr;
	logger.IsLogRotationEnabled = false;
	logger.LogsFolder = "/var/log/log-courier-rest"
	logger.LogFiles = []string{}
	logger.LogMinLevel = "INFO"
	logger.LogMaxLevel = "ERROR"
	logger.LogFormat = "%Date %Time [%LEVEL] %Msg%n"
	logger.LogRotationAttributes = LogRotationMetadata{RotationType: "size", MaxRollNumber: 5, Properties : map[string]string{}};
	
	if c, err := config.ReadDefault(logger.ConfigurationFileListPtr.GetLoggerConfFile()); err != nil {
		log.Errorf("unable to read lc-logger.conf file configPath:%s ", logger.ConfigurationFileListPtr.GetLoggerConfFile());
	}else {
		logger.LogConfigurationFile = logger.ConfigurationFileListPtr.GetLoggerConfFile();
		
		if temp, err := c.String("LC_LOGGER", "logrotation"); err == nil {
			if (strings.ToLower(strings.TrimSpace(temp)) == "enabled") {
				logger.IsLogRotationEnabled = true
			}else {
				logger.IsLogRotationEnabled = false;
			}
		}
		
		if temp, err := c.String("LC_LOGGER", "logfolder"); err == nil {
			logger.LogsFolder = temp;
		}
		
		if temp, err := c.String("LC_LOGGER", "minloglevel"); err == nil {
			logger.LogMinLevel = strings.ToUpper(strings.TrimSpace(temp));
		}
		
		if temp, err := c.String("LC_LOGGER", "maxloglevel"); err == nil {
			logger.LogMaxLevel = strings.ToUpper(strings.TrimSpace(temp));
		}
		
		if temp, err := c.String("LC_LOGGER", "logformat"); err == nil {
			logger.LogFormat = temp;
		}
		
		if temp, err := c.String("LC_LOGGER", "rotation_interval_type"); err == nil {
			logger.LogRotationAttributes.RotationType = strings.ToLower(strings.TrimSpace(temp));
		}
		
		if temp, err := c.String("LC_LOGGER", "max_rotated_files"); err == nil {
			if logger.LogRotationAttributes.MaxRollNumber,err = strconv.Atoi(temp) ; err != nil {
				logger.LogRotationAttributes.MaxRollNumber = 5;
				
			}
		}
		
		if (logger.LogRotationAttributes.RotationType == "size") {
			logger.LogRotationAttributes.Properties["rotation_size_limit_inbytes"] = "26214400"
			if temp, err := c.String("LC_LOGGER", "rotation_size_limit"); err == nil {
			   patternObj, _ := regexp.Compile(LogFileSizePattern);
			   if foundBool := patternObj.MatchString(temp); foundBool {
				  	sizeAttribs := patternObj.FindAllString(temp,-1);
				  	var bytes int = 0;
				  	var multiplier int = 1;
				  	var err error;
			  	    if bytes, err =  strconv.Atoi(sizeAttribs[0]); err != nil {
			  		       bytes = 26214400
			  	     }else {
			  	 	       multiplierStr := strings.ToUpper(sizeAttribs[1]);
			  			   switch multiplierStr {
			  					case "B"  : multiplier = 1;
			  					case "KB" : multiplier = 1024;
			  					case "MB" : multiplier = 1024*1024;
			  					case "GB" : multiplier = 1024*1024*1024;
			  					case "TB" : multiplier = 1024*1024*1024*1024;
			  					default:
			  			        		    multiplier = 1;
			  			        		    }
	  			     bytes = bytes * multiplier;
	  			     logger.LogRotationAttributes.Properties["rotation_size_limit_inbytes"] = strconv.Itoa(bytes);
			  	}
			  }
		    }
		}
	}
	
	logger.LogFiles = append(logger.LogFiles, logger.LogsFolder + "/" + logger.RelativeFileName + ".log");
	seelogStart := `<seelog minlevel="`+strings.ToLower(logger.LogMinLevel)+`" maxlevel="`+strings.ToLower(logger.LogMaxLevel)+`">`
	outputsStart := `<outputs formatid="common">`
	rollingSetting := ``;
	if (logger.IsLogRotationEnabled) {
		if (logger.LogRotationAttributes.RotationType == "size") {
			rollingSetting = `<rollingfile type="size" filename="`+logger.LogFiles[0]+`" maxsize="`+
			logger.LogRotationAttributes.Properties["rotation_size_limit_inbytes"]+`" maxrolls="`+strconv.Itoa(logger.LogRotationAttributes.MaxRollNumber)+`"/>`;
		}
	}
	formatsStart := `<formats>`
	format1 := `<format id="common" format="`+logger.LogFormat+`" />`
	formatsEnd := `</formats>`
	outputsEnd := `</outputs>`;
	seelogEnd := `</seelog>`;
	
	appconfig := seelogStart + outputsStart + rollingSetting +  outputsEnd + formatsStart + format1 + formatsEnd + seelogEnd;
	logger.Logger , _ = log.LoggerFromConfigAsBytes([]byte(appconfig));  
	      
}


func (logger *LoggerImpl) Info(message ...interface{}) {
	logger.Logger.Info(message);
	logger.Logger.Flush();
}

func (logger *LoggerImpl) Warn(message ...interface{}) {
	logger.Logger.Warn( message);
	logger.Logger.Flush();
}

func (logger *LoggerImpl) Debug(message ...interface{}) {
	logger.Logger.Debug(message);
	logger.Logger.Flush();
}

func (logger *LoggerImpl) Error(message ...interface{}) {
    logger.Logger.Error(message)
    logger.Logger.Flush();	
}

func (logger *LoggerImpl) Critical(message ...interface{}) {
	logger.Logger.Critical(message);
	logger.Logger.Flush();
}

func (logger *LoggerImpl) Trace(message ...interface{}) {
	logger.Logger.Trace(message);
	logger.Logger.Flush();
}