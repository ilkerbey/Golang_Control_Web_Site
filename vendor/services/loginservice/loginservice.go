package loginservice

import (
	"models"
	"services/database"
	"fmt"
	)

func  GetloginDB(Username ,Password string)models.Response{
	var logs models.User
	query := "SELECT ID FROM users WHERE Username = ? AND Password = ?"
	err := database.DB.QueryRow(query,Username,Password).Scan(&logs.ID)


	if err != nil{
		fmt.Println(err)


	}
	var response models.Response
	response.Message = "Giriş başarısız.Lütfen  tekrar deneyiniz."
	if logs.ID > 0 {
		response.Name = Username
		response.ID=logs.ID
		response.Status=true
		response.Message="Giriş Başarılı"
	}

	return response


}

func Register(user models.User)models.Response{

	query := "INSERT users SET Username = ?,Password = ?,Email = ? ,Adi = ?,Soyadi = ?"
	res,err := database.DB.Exec(query,user.Username,user.Password,user.Email,user.Adi,user.Soyadi)

	var response models.Response

	if err !=nil{
		fmt.Println(err)
		response.Message="Eklenirken Bir Hata Oluştu."
		return response
	}

	response.Status=true
	response.Message= "Başarı ile Eklendi"
	lastid,_:=res.LastInsertId()
	response.ID=int(lastid)

	return response
}

func GetUserInfo(ID int)models.User{
	var user models.User
	query := "SELECT ID, Username,Email,Adi,Soyadi FROM users WHERE ID = ?"
	err := database.DB.QueryRow(query,ID).Scan(&user.ID,&user.Username,&user.Email,&user.Adi,&user.Soyadi)

	if err != nil {
		fmt.Println(err)
	}

	return user

}

func GetUsernameControl(username string)bool{
	var user models.User
	query := "SELECT ID From users WHERE Username = ?"
	err := database.DB.QueryRow(query,username).Scan(&user.ID)

	if err != nil {
		fmt.Println(err)
	}

	if user.ID >0 {
		return false
	}

	return true
}
func GetEmailControl(email string)bool{
	var user models.User
	query := "SELECT ID From users WHERE Email = ?"
	err := database.DB.QueryRow(query,email).Scan(&user.ID)

	fmt.Println()

	if err != nil {
		fmt.Println(err)
	}

	if user.ID> 0 {
		return false
	}

	return true
}
