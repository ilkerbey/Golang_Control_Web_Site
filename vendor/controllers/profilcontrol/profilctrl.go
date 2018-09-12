package profilcontrol

import (
	"net/http"
	"models"
	"services/layoutservice"
	"services/display"
	"github.com/gorilla/mux"
		"strconv"
	"fmt"
	"services/loginservice"
	"services/hmacservice"
)

func Index(w http.ResponseWriter,r *http.Request) {
	var data models.Home
	vars := mux.Vars(r)
	IDstr := vars["ID"]
	ID, err := strconv.Atoi(IDstr)
	if err != nil {
		fmt.Println(err)
	}


	data.Page=layoutservice.GetSharedData("profil","profil",r)
	data.User=loginservice.GetUserInfo(ID)
	fmt.Println(data.User.Username)


	display.View(w,r,"profilIndex",data)
}

func Profil(w http.ResponseWriter, r *http.Request){
	var data models.Home
	var user models.User
	var id = hmacservice.GetCurrentUser(r).ID
	user = hmacservice.GetCurrentUser(r)
	fmt.Println("sss",user.ID)
	data.Page = layoutservice.GetSharedData("Profilim","Profil sayfam",r)
	data.User = loginservice.GetUserInfo(id)
	display.View(w,r,"profilIndex",data)
}

func Update(w http.ResponseWriter,r *http.Request){

}