package homectrl

import (
	"net/http"
	"models"
	"services/display"

	"services/menuservice"
	"strings"
	"github.com/nu7hatch/gouuid"
	"services/layoutservice"
	"services/photoservice"
	"fmt"
	"services/hmacservice"
	"services/loginservice"
)

func Index(w http.ResponseWriter,r *http.Request){
	var data models.Home
	data.Page = layoutservice.GetSharedData("Anasayfa","Anasayfa",r)

	var ID = hmacservice.GetCurrentUser(r).ID
	data.User = loginservice.GetUserInfo(ID)
	fmt.Println(data.User.ID)

	display.View(w,r,"homeIndex",data)
}
func GetCookiew(w http.ResponseWriter,r *http.Request) *http.Cookie{
	c,err := r.Cookie("session")

	if err != nil {
		sID,_ := uuid.NewV4()
		c = &http.Cookie{
			Name:"session",
			Value:sID.String(),
		}
		http.SetCookie(w,c)
	}
	return c
}

func IndexPost(w http.ResponseWriter,r *http.Request){

	var data models.Home

	data.Page.Category = menuservice.GetMenuList()

	var model models.Photo
	c := GetCookiew(w,r)

	fmt.Println("Index Post Buradasın")
	model = photoservice.NormalPhotoSave(r)
	fmt.Println("Çıktı Başarılı Geri DÖndü. ")

	c =AppendValue(w,c,model.Pname.String)
	data.Status=model.Status
	//data.FileNames = strings.Split(c.Value,"|")
	data.FileName = model.Pname.String
	display.View(w,r,"homeIndex",data)


}


func AppendValue(w http.ResponseWriter, cookie *http.Cookie,fname string)*http.Cookie{

	s := cookie.Value

	if !strings.Contains(s,fname){
		s += "|" + fname
	}

	cookie.Value = s
	http.SetCookie(w,cookie)

	return cookie
}