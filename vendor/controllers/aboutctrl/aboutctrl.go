package aboutctrl

import (
	"net/http"
	"models"
	"services/layoutservice"
	"services/display"
	"services/aboutservice"
	"github.com/gorilla/mux"
	"fmt"
)

func Index(w http.ResponseWriter , r *http.Request) {

	vars := mux.Vars(r)
	Constr := vars["PageSlug"]
	var data models.Custom
	data.PageInfo = aboutservice.GetPageInfo(Constr)
	if data.PageInfo.ID == 0{
		fmt.Println("Buradan Yönlendirdi ....")
		http.Redirect(w,r,"http://localhost:8080",http.StatusSeeOther)
		return
	}
	data.Page = layoutservice.GetSharedData("Hakkında",data.PageInfo.Title,r)
	display.View(w,r,"aboutIndex",data)
}
