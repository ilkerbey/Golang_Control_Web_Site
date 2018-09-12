package signupctrl

import (
	"models"
	"net/http"
	"services/menuservice"
	"services/display"
	"services/loginservice"
	"fmt"
	"services/hmacservice"
)

func Index(w http.ResponseWriter,r *http.Request) {
	var data models.Home

	data.Page.Title = "Kayıt ol"
	data.Page.Description = "Kayıt ol"
	data.Page.Category = menuservice.GetMenuList()

	display.View(w,r,"signupIndex",data)
}

func IndexPost(w http.ResponseWriter,r *http.Request){

	var data models.Home

	data.PostOlduMu = true

	data.Page.Title = "Kayıt ol"
	data.Page.Description = "Kayıt ol"
	data.Page.Category = menuservice.GetMenuList()

	var signup models.User




	var username = r.FormValue("username")
	KullaniciVarmi := loginservice.GetUsernameControl(username)
	fmt.Println(KullaniciVarmi)
	if KullaniciVarmi == false {
		data.Response.Message="Böyle Bir kullanıcı Var"
		display.View(w, r, "signupIndex", data)
		return

	}
	signup.Username=r.FormValue("username")
	var as = r.FormValue("password")
	var al = r.FormValue("password_2")
	if as != al {
		data.Response.Message = "Lütfen Şifreniz Aynı olsun"
		display.View(w, r, "signupIndex", data)
		return
	}


	signup.Password=r.FormValue("password")


	var email = r.FormValue("email")
	Emailvarmi := loginservice.GetEmailControl(email)
	if Emailvarmi == false {
		fmt.Println("selam")
		data.Response.Message="Bu email Mevcuttur. Lütfen Başka Bir Emaili deneyiniz."
		display.View(w,r,  "signupIndex", data)
		return
	}
	signup.Email=r.FormValue("email")

	signup.Adi=r.FormValue("name")
	signup.Soyadi=r.FormValue("lastname")

	if r.Method==http.MethodPost{

		var sifre = hmacservice.RegisterSifrele(signup.Password)


		signup.Password=sifre

		fmt.Println("Yeni oluşan şifre",signup.Password)


		data.Response = loginservice.Register(signup)
		fmt.Println("Kayıt başarılı.")
		sifree := hmacservice.RegisterSifreCoz(sifre)
		fmt.Println("Şifremizin Gerçek Hali",sifree.Password)



	}

	display.View(w,r,"signupIndex",data)

}
