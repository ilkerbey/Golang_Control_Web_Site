package login

import (
	"net/http"
	"models"
	"services/display"
	"fmt"
	"services/hmacservice"
	"services/layoutservice"
	"services/loginservice"
)

func Index(w http.ResponseWriter,r *http.Request){
	fmt.Println("LoginIndex  = ",hmacservice.IsAuth(r))
	var data models.Home
	data.Page = layoutservice.GetSharedData("Login","login",r)

	var log models.User

	log.Username = r.FormValue("username")
	log.Password= r.FormValue("password")





	display.View(w,r,"loginIndex",data)
}

func IndexPost(w http.ResponseWriter,r *http.Request){
	var data models.Home
	data.PostOlduMu = true
	data.Page = layoutservice.GetSharedData("Login","Login",r)

	var log models.User

	log.Username = r.FormValue("username")
	log.Password=r.FormValue("password")
	if r.Method == http.MethodPost{

		var sifreCoz = hmacservice.RegisterSifrele(log.Password)
		log.Password=sifreCoz
		data.Response = loginservice.GetloginDB(log.Username,log.Password)

		if data.Response.Status == true{
			data.Page.GirisYapmismi=true
			var authCode = hmacservice.Sifrele(data.Response.ID,data.Response.Name)
			fmt.Println(authCode)
			c := hmacservice.SetCookieHmac(w,r,authCode)
			fmt.Println(c)
			http.Redirect(w,r,"/anasayfa",http.StatusSeeOther)
		}
	}


	display.View(w,r,"loginIndex",data)

}
