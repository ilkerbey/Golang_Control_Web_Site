package layoutservice

import (
	"models"
	"services/hmacservice"
	"net/http"
	"services/menuservice"
	"html/template"
)

func GetSharedData(title,description string ,r *http.Request)models.SimplePage{
	var data models.SimplePage
	data.GirisYapmismi = hmacservice.IsAuth(r)
	data.Category = menuservice.GetMenuList()
	data.Title = title
	data.Description=description
	return data
}
func GetSharedDataAboout(title string,description template.HTML ,r *http.Request)models.SimplePage{
	var data models.SimplePage
	data.GirisYapmismi = hmacservice.IsAuth(r)
	data.Category = menuservice.GetMenuList()
	data.Title = title
	data.Abouts=description
	return data
}