package photoservice

import (
	"models"
	"services/database"
	"fmt"
	"services/convert"
		"strings"
	"crypto/sha1"
	"io"
	"os"
	"path/filepath"
	"net/http"
	"services/hmacservice"
			"log"
	"github.com/disintegration/imaging"
)

func Insert(model models.Photo)models.Response{
	query := "INSERT photo SET PhotoName=?,UserID = ?, Date=NOW()"
	res,err :=database.DB.Exec(query,model.Pname,model.UserInfo.ID)
	fmt.Println("Adı : ",model.Pname,"Id'si",model.UserInfo.ID)
	fmt.Println("İnsert bölümü")
	var response models.Response


	fmt.Println("i-1")

	if err != nil{
		response.Message = "Bir hata oluştu.Lütfen tekrar deneyiniz."
		fmt.Println(err)
		return  response
	}

	fmt.Println("i - 2")
	response.Status =true
	response.Message = "İşlem başarılı.Fotoğraf başarı ile kayıt edildi."
	lastid,_ :=res.LastInsertId()
	response.ID = int(lastid)
	return response
}



func DeleteInfo(ID int)models.Response{
	query := "DELETE FROM photo WHERE ID = ?"
	_ ,err := database.DB.Exec(query,ID)

	var response models.Response

	if err != nil {
		fmt.Println(err)
		response.Message = "Silinirken Bir Hata Oluştu"
		return response
	}

	response.Status = true
	response.Message = "Kayıt Başarılı ile  silindi."

	return response
}

func GetList(userID,listType int) []models.Photo{
	var args []interface{}
	query := "SELECT u.Username, b.ID , b.PhotoName , b.Date FROM photo b LEFT JOIN users u ON b.UserID = u.ID "
	if listType == 0{
		query += "WHERE UserID=? "
		args = append(args,userID)
	}
	query += "ORDER BY ID DESC"

	rows,err := database.DB.Query(query,args...)
	defer rows.Close()
	if err != nil{
		fmt.Println(err)
	}
	var list []models.Photo
	for rows.Next(){
		var model models.Photo
		rows.Scan(&model.UserInfo.Username,&model.ID,&model.Pname,&model.Date)
		model.Datestr = convert.ToDateString(model.Date)
		list =append(list,model)
	}
	return list
}

func GetPhotoInfo(ID int)models.Photo{

	var model models.Photo
	query := "SELECT ID,PhotoName,Date FROM photo WHERE ID = ?"
	err := database.DB.QueryRow(query,ID).Scan(&model.ID,&model.Pname,&model.Date)

	if err != nil {
		fmt.Println(err)
	}

	return model
 }

func NormalPhotoSave(r *http.Request)models.Photo{
	var model models.Photo

	model.UserInfo = hmacservice.GetCurrentUser(r)
	mf,fh,err := r.FormFile("nf")

	fmt.Println("1")

	if err != nil {
		fmt.Println(err)
	}
	defer mf.Close()

	ext := strings.Split(fh.Filename,".")[1]
	h := sha1.New()
	io.Copy(h,mf)

	model.Pname.Scan(fmt.Sprintf("%x",h.Sum(nil))+"."+ext)
	fmt.Println(model.Pname.String)

	wd,err := os.Getwd()


	fmt.Println("2")

	if err != nil {
		fmt.Println(err)
	}

	path := filepath.Join(wd,"public","pics",model.Pname.String)
	nf,err := os.Create(path)

	fmt.Println("3")
	if err != nil {
		fmt.Println(err)
	}
	var response models.Response
	response = Insert(model)
	model.Status = response.Status


	defer nf.Close()

	mf.Seek(0,0)
	io.Copy(nf,mf)

	SmallPhotoSave(model.Pname.String)

	return model
}

func SmallPhotoSave(resimadi string){
	//Resim Dosyamızı ismi ile açıyoruz.
	file,err := imaging.Open("public/pics/"+resimadi)

	if err != nil {
		fmt.Println(err)
	}


	//resim dosyamızın doyutunu 200px ile boyutlandırıyoruz.
	src := imaging.Resize(file,200,200,imaging.Lanczos)


	err = imaging.Save(src,"public/pics/tmpl_"+resimadi)

	if err != nil {
		log.Fatalf("Failed to save image: %v",err)
	}




}