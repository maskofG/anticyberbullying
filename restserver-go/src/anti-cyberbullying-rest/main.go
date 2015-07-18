/**
* Copyright (c) 2015 EMC Corporation
* All Rights Reserved
*
* This software contains the intellectual property of EMC Corporation
* or is licensed to EMC Corporation from third parties.  Use of this
* software and the intellectual property contained therein is expressly
* limited to the terms and conditions of the License Agreement under which
* it is provided by or on behalf of EMC.
*
**/

/**
*
* LogCourier REST Service main module
*
**/
package main

import (
	"flag"
	"path/filepath"
	"os"
)  







var ConfigurationFilesImplPtr ConfigurationFilesInterface;
var StorageBackendImplPtr     StorageBackendInterface;
var RestApiImplPtr            CyberBullyingEntryPointRestApiInterface;
var RestServerImplPtr         RestServerInterface;
//main it calls all the init functions and starts the http server to REST service
func main() {
	ConfigurationFilesImplPtr = new(ConfigurationFilesImpl);
	ArgsInit();
	StorageBackendImplPtr = new(ElasticsearchStorageBackendImp);
	StorageBackendImplPtr.init(ConfigurationFilesImplPtr);
	RestApiImplPtr = new(CyberBullyingEntryPointRestApiImpl);
	RestApiImplPtr.init(ConfigurationFilesImplPtr, StorageBackendImplPtr);
	RestServerImplPtr = new(RestServer);
	RestServerImplPtr.init(ConfigurationFilesImplPtr, RestApiImplPtr.GetApi());
    Run(); 
}


func ArgsInit(){
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	RestConfigPath := flag.String("rest-conf", dir + "/" + "../src/conf/rest-conf.conf", "lcrest configuration absolute path")
	ESConfigPath := flag.String("es-conf", dir + "/" + "../src/conf/es-conf.conf", "log courier rest logger configuration");
	LoggerConfigPath := flag.String("logger-conf", dir + "/" + "../src/conf/logger-conf.conf", "log courier rest logger configuration");
	flag.Parse();
	
	if filepath.IsAbs(*RestConfigPath) == false {
		*RestConfigPath,_ = filepath.Abs(*RestConfigPath)
	}
	
	if filepath.IsAbs(*LoggerConfigPath) == false {
		*LoggerConfigPath,_ = filepath.Abs(*LoggerConfigPath)
	}
	
	if filepath.IsAbs(*ESConfigPath) == false {
		*ESConfigPath,_ = filepath.Abs(*ESConfigPath)
	}
	
	ConfigurationFilesImplPtr.SetRestConfFile(*RestConfigPath);
	ConfigurationFilesImplPtr.SetESConfFile(*ESConfigPath);
	ConfigurationFilesImplPtr.SetLoggerConfFile(*LoggerConfigPath);
}


//Run starts REST service
func Run(){
	RestApiImplPtr.RegisterApi()
	RestServerImplPtr.StartRestServer();
}

