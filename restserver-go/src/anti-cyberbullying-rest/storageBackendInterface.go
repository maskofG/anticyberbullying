package main;


type StorageBackendInterface interface {
	init(ConfigurationFilesImplPtr ConfigurationFilesInterface)
    AddDataById(id string, data string)(err error)
    AddData(data string)(id string, err error)
    GetData(id string) (data string, err error)
    GetAllData()(data string, err error)
    DeleteData(id string) (err error)
    UpdateData(id string, data string)(err error)
}