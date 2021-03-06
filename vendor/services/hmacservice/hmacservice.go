package hmacservice

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
	"net/http"
	"models"
	"services/convert"
	"services/cookieservice"
)

func Sifrele(ID int,username string) string{
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = jwt.MapClaims{
		"uName": username,
		"uyeID":ID,
		"exp":time.Now().Add(time.Duration(48)*time.Hour).Unix(),
	}
	tokenString,err := token.SignedString([]byte("emanuel"))

	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("Oluşturulan Uye Cookie Şifremiz",tokenString)
	return tokenString
}

func RegisterSifrele(password string) string{
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = jwt.MapClaims{
		"uPass": password,
	}

	tokenstring,err := token.SignedString([]byte("sifre"))
	if err!= nil{
		fmt.Println("Register Sifrele Hata : ",err)
	}
	fmt.Println("Oluşturulan Register Sifremiz : ",tokenstring)
	return tokenstring
}


func SifreyiCoz(cozulecekSifre string)(userData models.User){
	token, err := jwt.Parse(cozulecekSifre, func(token *jwt.Token) (interface{}, error) {
		return []byte("emanuel"), nil
	})

	if err == nil && token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//Burada hata yok olmadığını gösterir.Claimleri Çözdü
			var email = claims["uName"]
			var uyeID = claims["uyeID"]
			var strEmail = fmt.Sprintf("%v",email)
			var strUyeID = fmt.Sprintf("%v",uyeID)
			userData.ID = convert.ToInt(strUyeID)
			userData.Username = strEmail
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println(err)
	return

}


func RegisterSifreCoz(cozuleceksifre string)(userdata models.User){
		token , err := jwt.Parse(cozuleceksifre, func(token *jwt.Token) (interface{}, error) {
			return []byte("sifre"),nil
		})

		if err == nil && token.Valid{
			if claims,ok := token.Claims.(jwt.MapClaims);ok && token.Valid{
				var uPass = claims["uPass"]
				var strEmail = fmt.Sprintf("%v",uPass)

				userdata.Password=strEmail
				return
			}
			return
		}
return
}

func  	SetCookieHmac(w http.ResponseWriter,r *http.Request,sifre string )*http.Cookie{
	c , _ := r.Cookie("authCode")
		c  = &http.Cookie{
			Name:"authCode",
			Value:sifre,
			Expires:time.Now().Add(time.Hour*48),
		}


		http.SetCookie(w,c)

		return c
}

func IsAuth (r *http.Request) bool{

	c,_:= r.Cookie("authCode")
	fmt.Println("Cookie Şifremiz : ",c)
	if c == nil{
		return false
	}
	var userData = SifreyiCoz(c.Value)
	if userData.ID > 0{
	return true
	}
	return false
}

func GetCurrentUser(r *http.Request) models.User{
	var authCode = cookieservice.GetCookieValue(r,"authCode")
	var userInfo = SifreyiCoz(authCode)
	return userInfo
}
