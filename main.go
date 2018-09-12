package main

import (
	"services/display"
	"net/http"
	"github.com/gorilla/mux"
	"controllers/homectrl"
	"services/database"
	_ "github.com/go-sql-driver/mysql"
	"controllers/photoctrl"
	"controllers/login"
	"fmt"
	"services/hmacservice"
	"controllers/signupctrl"
	"controllers/aboutctrl"
	"controllers/profilcontrol"
)

func securedPage(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("*****************111111111111111111111*****************")
		var girisYapmisMi = hmacservice.IsAuth(r)
		fmt.Println("******************2222222222222222222************")
		if girisYapmisMi {
			f(w, r)
			return
		}
		fmt.Println("burdayım")
		http.Redirect(w,r,"http://localhost:8080/login",http.StatusSeeOther)
	}
}


func main(){

	display.LoadTemplates()

	database.Connect("mysql","golang")
	go fmt.Println("Program Çalışıyor ..... ")

	defer database.DB.Close()

	var r = mux.NewRouter()



	r.HandleFunc("/",homectrl.Index )
	r.HandleFunc("/editpic",securedPage(photoctrl.Index))
	r.HandleFunc("/editpic/duzenle/{ID}",photoctrl.Update).Methods("GET")
	r.HandleFunc("/editpic/duzenle/{ID}",photoctrl.UpdatePost).Methods("POST")
	r.HandleFunc("/editpic/delete",photoctrl.DeletePhoto)
	r.HandleFunc("/anasayfa", securedPage(homectrl.Index )).Methods("GET")
	r.HandleFunc("/anasayfa", securedPage(homectrl.IndexPost)).Methods("POST")
	r.HandleFunc("/profil",securedPage(profilcontrol.Profil))
	r.HandleFunc("/profil/{ID}",securedPage(profilcontrol.Index))
	r.HandleFunc("/login",login.Index).Methods("GET")
	r.HandleFunc("/login",login.IndexPost).Methods("POST" )
	r.HandleFunc("/cikis",login.LogoutIndex)
	r.HandleFunc("/signup",signupctrl.Index).Methods("GET")
	r.HandleFunc("/signup",signupctrl.IndexPost).Methods("POST")
	r.HandleFunc("/{PageSlug}",aboutctrl.Index)


	http.Handle("/",r)

	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
	http.Handle("/public/",http.StripPrefix("/public/",http.FileServer(http.Dir("public"))))


	http.ListenAndServe(":8080",nil)
}
