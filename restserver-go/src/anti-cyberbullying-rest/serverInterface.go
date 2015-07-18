package main


import "github.com/ant0ine/go-json-rest/rest"

type RestServerInterface interface {
	init(ConfigurationFileListPtr ConfigurationFilesInterface, Api *rest.Api);
	StartRestServer()
}