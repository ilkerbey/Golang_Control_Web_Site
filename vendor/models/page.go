package models

import "html/template"

type SimplePage struct {
	Title string
	Description string
	Abouts template.HTML
	Category []Part
	GirisYapmismi bool
}

type Home struct {
	Page SimplePage
	FileName string
	Status bool
	Response Response
	PostOlduMu bool
	User User

}
type Part struct {
	Category string
	PageSlug string
}

type PhotoPage struct {
	Page SimplePage
	ListType int
	PhotoList []Photo
}

type PhotoUpdate struct {
	Page SimplePage
	PhotoInfo Photo
}
type PhotoDelete struct {
	Page SimplePage
	PhotoInfo Photo
	Response Response
}

type Custom struct {
	Page SimplePage
	PageInfo PageInfo
}

type PageInfo struct{
	ID int
	PageSlug string
	Title string
	Content string
	ContentHTML template.HTML
}

