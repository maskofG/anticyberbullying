package main;


import (
	"strings"
	"encoding/hex"
)


var BullyTerms []string = []string {"asshole", "bitches" , "fuckyou" , "getsomelife", "motherfucker"}

type SignUpUserData struct {
	Name               string  `json:"name"`
	EmailId            string  `json:"emailid"`
	Password           string  `json:"password"`
}

type PersonalInfo struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	EmaildId      string `json:"emailid"`
	Password      string `json:"password"`
}

type ChildInfo struct {
	Id                             string            `json:"id"`
	ChildMinimalInformation        ChildMinimalInfo  `json:"minimalinfo"`
	ChildExtendedInformation       ChildExtendedInfo `json:"extendedinfo"`
	OSNHandlesInformation          []OSNHandleInfo
}

type ChildMinimalInfo struct {
	Name                          string  `json:"name"`
	Gender                        string  `json:"gender"`
	Age                           string  `json:"age"`
}

type ChildExtendedInfo struct {
	Birthday               string  `json:"birthdate"`
	Hobbies                []string    `json:"hobbies"`
	SchoolName             string      `json:"schoolname"`
	CollegeName            string      `json:"collegename"`
}

type OSNHandleInfo struct {
	OSNName               string   `json:"osnName"`   //Initially facebook or twitter
	Handle                string   `json:"handle"`    //OSN profile identifier
}

type UserData  struct {
   PersonalInformation  PersonalInfo `json:"personalInformation"`
   ChildInforamtion     []ChildInfo  `json:"chilrenDetails"`
}

func (ud  *UserData) GetUserId() (userid string) {
	userid = "";
	userid = ud.PersonalInformation.Id
	return userid;
}

func (ud  *UserData) toString() (serializeddata string) {
	serializeddata = "";
	
	return serializeddata;
}

func (ud  *UserData) createChildId(childname string) (childid string) {
	childid = "";
	childid = strings.TrimSpace(ud.PersonalInformation.EmaildId) + "_";
	childname = strings.TrimSpace(childname);
	terms := strings.Split(childname, " ");
	for index, term := range terms {
		term = strings.TrimSpace(term);
		terms[index] = term;
	}
	childid = childid + strings.Join(terms, "_");
	childid = hex.EncodeToString([] byte(childid));
	return childid;
}


func (ud  *UserData) GetChildId(childname string) (childid string) {
	childid = "";
	for index, childinfo := range ud.ChildInforamtion {
		if (childinfo.ChildMinimalInformation.Name == childname) {
			if (childinfo.Id == "" ) {
				childinfo.Id = ud.createChildId(childname);
				ud.ChildInforamtion[index] = childinfo
				break;
			}
			childid = childinfo.Id;
		}
	}
	return childid;
}









