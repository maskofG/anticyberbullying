/**
*
* EXPOSING REST API 
*
**/

package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"encoding/json"
	"strings"
	"crypto/md5"
    "encoding/hex"
)

type CyberBullyingEntryPointRestApiImpl struct {
	LoggerImplPtr              LoggerInterface;
	Api                        *rest.Api;
	StorageBackendImplPtr      StorageBackendInterface;
}


func (restapi *CyberBullyingEntryPointRestApiImpl) init(ConfigurationFilesImplPtr ConfigurationFilesInterface, StorageBackendImplPtr StorageBackendInterface){
    restapi.LoggerImplPtr = new(LoggerImpl);
    restapi.LoggerImplPtr.init(ConfigurationFilesImplPtr, "anticyberbullying-entrypoint-restapi");
    restapi.SetApi(nil);
    restapi.StorageBackendImplPtr = StorageBackendImplPtr
}

func (restapi *CyberBullyingEntryPointRestApiImpl) GetApi() (Api *rest.Api) {
	if (restapi.Api == nil) {
		restapi.Api = rest.NewApi()
	 	restapi.Api.Use(rest.DefaultDevStack...)
	}
	return restapi.Api;
}


func (restapi *CyberBullyingEntryPointRestApiImpl) SetApi(Api *rest.Api){
	if (Api == nil) {
		restapi.Api = rest.NewApi()
		restapi.Api.Use(rest.DefaultDevStack...)
		restapi.Api.Use(&rest.CorsMiddleware{
        						RejectNonCorsRequests: false,
        						OriginValidator: func(origin string, request *rest.Request) bool {
            													return true // TODO once deploy on AWS
        										},
        						AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
                                AllowedHeaders: []string{
            									"Accept", "Content-Type", "X-Custom-Header", "Origin"},
        						AccessControlAllowCredentials: true,
        						AccessControlMaxAge:           3600,
    					})
	}else {
		restapi.Api = Api;
	}
}


func (restapi *CyberBullyingEntryPointRestApiImpl) RegisterApi(){
	 if restapi.Api == nil {
	 	restapi.SetApi(nil);
	 	}
     
     router, err := rest.MakeRouter(
     	rest.Post("/api/signup", restapi.SignUpUser),
     	rest.Post("/api/signin", restapi.SignUpUser),
     	rest.Get("/api/#parentid/childinfo", restapi.GetChildInfo),
     	rest.Post("/api/#parentid/childinfo", restapi.AddChildInfo),
     	rest.Put("/api/#parentid/childinfo", restapi.UpdateChildInfo),
     	rest.Get("/api/childbullydata/#parentid/#childid", restapi.GetChildBullyData),
     	rest.Post("/api/childbullydata/#parentid/#childid", restapi.PostChildBullyData),
     	rest.Post("/api/predict", restapi.PredictBullyData),
    )
    if err != nil {
        restapi.LoggerImplPtr.Error(err.Error());
        panic(err);
    }
    restapi.Api.SetApp(router)
}


func (restapi *CyberBullyingEntryPointRestApiImpl) SignUpUser(w rest.ResponseWriter, r *rest.Request) {
	var signupdata SignUpUserData;
	err := r.DecodeJsonPayload(&signupdata)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	isAlreadyPresent := false;
	if alreadyexistingdata, err := restapi.StorageBackendImplPtr.GetData(strings.TrimSpace(signupdata.EmailId)); err == nil {
		var tempUserData UserData;
        if err := json.Unmarshal([]byte(alreadyexistingdata), &tempUserData); err == nil {
        	if (tempUserData.PersonalInformation.EmaildId == signupdata.EmailId) {
        		isAlreadyPresent = true
        	}
        }
	}
	
	if (isAlreadyPresent == false) {
		var tempUserData UserData;
		tempUserData.PersonalInformation = PersonalInfo{};
		tempUserData.PersonalInformation.Name = signupdata.Name;
		tempUserData.PersonalInformation.EmaildId = signupdata.EmailId;
		tempUserData.PersonalInformation.Id = strings.TrimSpace(signupdata.EmailId);
		hash := md5.Sum([]byte(signupdata.Password))
		tempUserData.PersonalInformation.Password = hex.EncodeToString(hash[:])
		restapi.StorageBackendImplPtr.AddDataById(tempUserData.GetUserId(), tempUserData.toString());
	}else {
		rest.Error(w, "user already present" , http.StatusConflict);
	}
}

func (restapi *CyberBullyingEntryPointRestApiImpl) SignInUser(w rest.ResponseWriter, r *rest.Request) {
	    w.WriteJson(map[string]string{"Authenticated" : "true"});
}


func (restapi *CyberBullyingEntryPointRestApiImpl) GetChildInfo(w rest.ResponseWriter, r *rest.Request) {
	   w.WriteJson(map[string]string{"status" : "work still in progress"});
}

func (restapi *CyberBullyingEntryPointRestApiImpl) AddChildInfo(w rest.ResponseWriter, r *rest.Request) {
	    w.WriteJson(map[string]string{"status" : "work still in progress"});
}

func (restapi *CyberBullyingEntryPointRestApiImpl) UpdateChildInfo(w rest.ResponseWriter, r *rest.Request) {
	    w.WriteJson(map[string]string{"status" : "work still in progress"});
}

func (restapi *CyberBullyingEntryPointRestApiImpl) GetChildBullyData(w rest.ResponseWriter, r *rest.Request) {
	    w.WriteJson(map[string]string{"status" : "work still in progress"});
}


func (restapi *CyberBullyingEntryPointRestApiImpl) PostChildBullyData(w rest.ResponseWriter, r *rest.Request) {
	    w.WriteJson(map[string]string{"status" : "work still in progress"});
}


func (restapi *CyberBullyingEntryPointRestApiImpl) PredictBullyData(w rest.ResponseWriter, r *rest.Request) {
	    var payload []string = []string{}
	    if err := r.DecodeJsonPayload(&payload); err != nil {
	    	w.WriteJson("error parsing data");
	    	return;
	    }
	    isBully := false;
	    data := payload[0]
	    terms := strings.Split(data, " ");
	    for _, term := range terms {
	    	term = strings.TrimSpace(term);
	    	for _, bullyterm := range BullyTerms {
	    		if (bullyterm == term) {
	    			isBully = true;
	    		}
	    	}
	    }
	    
	    if (isBully == true) {
	    	w.WriteJson(map[string]string{"comment-type": "bully"});
	    }else {
	    	w.WriteJson(map[string]string{"comment-type": "non-bully"});
	    }
	    
}
