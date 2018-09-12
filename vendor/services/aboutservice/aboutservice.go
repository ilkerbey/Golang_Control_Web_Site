package aboutservice

import (
	"services/database"
	"fmt"
	"models"
	"html/template"
)

func GetPageInfo(pageSlug string) (models.PageInfo){
	query := "SELECT ID,Title,Content,PageSlug FROM Pages where PageSlug = ?"

	var model models.PageInfo

	row , err := database.DB.Query(query,pageSlug)
	if err != nil {
		fmt.Println(err)
	}

	defer row.Close()
	if row.Next(){
		err = row.Scan(&model.ID,&model.Title,&model.Content,&model.PageSlug)
	}

	model.ContentHTML= template.HTML(model.Content)



	return model

}
