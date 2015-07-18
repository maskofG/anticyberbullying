package main;

import 	"github.com/revel/config"


type ElasticsearchStorageBackendImp struct {
	LoggerImplPtr             LoggerInterface
	ConfigurationFilesImplPtr ConfigurationFilesInterface
	BaseUrl                   string
	Username                  string
	Password                  string
}


func (esb *ElasticsearchStorageBackendImp) init(ConfigurationFilesImplPtr ConfigurationFilesInterface) {
	esb.ConfigurationFilesImplPtr = ConfigurationFilesImplPtr;
	esb.LoggerImplPtr = new(LoggerImpl);
	esb.LoggerImplPtr.init(ConfigurationFilesImplPtr, "elasticsearch-storage-backend");
	esb.readESConfFile();
}

func (esb *ElasticsearchStorageBackendImp) readESConfFile() {
	if c, err := config.ReadDefault(esb.ConfigurationFilesImplPtr.GetESConfFile()); err != nil {
		esb.LoggerImplPtr.Error("unable to read es-conf.conf file switching to defualt");
		esb.BaseUrl = "http://0.0.0.0:9200"
	}else {	
		var protocol string="http"
		var ipaddress string="0.0.0.0"
		var port   string="9200"
		if temp, err := c.String("ES_SERVER", "protocol"); err == nil {
			protocol = temp;
		}
		
		if temp, err := c.String("ES_SERVER", "server"); err == nil && temp != "ES_ENDPOINT"{
			ipaddress = temp;
		}
		
		if temp, err := c.String("ES_SERVER", "restport"); err == nil && temp != "ES_REST_PORT" {
			port = temp;
		}
		
		if temp, err := c.String("ES_SERVER", "username"); err == nil {
			esb.Username = temp;
		}
		
		if temp, err := c.String("ES_SERVER", "password"); err == nil {
			esb.Password = temp;
		}
		
		esb.BaseUrl = protocol + "://" + ipaddress + ":"  + port;
	}
}

func (esb *ElasticsearchStorageBackendImp) connect() {
	
}


func (esb *ElasticsearchStorageBackendImp) AddDataById(id string, data string)(err error){
	err = nil;
	
	return err;
}


func (esb *ElasticsearchStorageBackendImp) AddData(data string)(id string, err error){
	id = ""
	err = nil;
	
	return id, err;
}


func (esb *ElasticsearchStorageBackendImp) GetData(id string) (data string, err error){
	data = "";
	err = nil;
	
	return data, err;
}


func (esb *ElasticsearchStorageBackendImp) GetAllData()(data string, err error){
	data = "";
	err = nil;
	
	
	return data, err;
}


func (esb *ElasticsearchStorageBackendImp) DeleteData(id string) (err error){
	err = nil
	
	return err;
}


func (esb *ElasticsearchStorageBackendImp) UpdateData(id string, data string)(err error){
	err = nil;
	
	return err;
}
