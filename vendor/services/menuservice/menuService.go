package menuservice

import "models"

func GetMenuList()(list []models.Part){
	var data models.Part
	data.Category = "Ana sayfa"
	data.PageSlug="anasayfa"
	list = append(list,data)
	data.Category = "Edit pic"
	data.PageSlug="editpic"
	list = append(list,data)
	data.Category = "Profil"
	data.PageSlug="profil"
	list = append(list,data)
	data.Category = "Hakkımızda"
	data.PageSlug="hakkimizda"
	list = append(list,data)
	data.Category = "İletişim"
	data.PageSlug="iletisim"
	list = append(list,data)
	return list
}