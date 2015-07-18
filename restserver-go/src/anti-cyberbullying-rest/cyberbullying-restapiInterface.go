package main

import "github.com/ant0ine/go-json-rest/rest"

type CyberBullyingEntryPointRestApiInterface interface {
	init(ConfigurationFilesImplPtr ConfigurationFilesInterface,  StorageBackendImplPtr StorageBackendInterface)
    GetApi() (Api *rest.Api)
    SetApi(Api *rest.Api)
    RegisterApi()
    SignUpUser(w rest.ResponseWriter, r *rest.Request)
    SignInUser(w rest.ResponseWriter, r *rest.Request)
    GetChildInfo(w rest.ResponseWriter, r *rest.Request)
    AddChildInfo(w rest.ResponseWriter, r *rest.Request)
    UpdateChildInfo(w rest.ResponseWriter, r *rest.Request)
    GetChildBullyData(w rest.ResponseWriter, r *rest.Request)
    PostChildBullyData(w rest.ResponseWriter, r *rest.Request)
}
