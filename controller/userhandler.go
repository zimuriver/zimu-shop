package controller

import (
	"log"
	"net/http"
	"text/template"
	"zimu-shop/utils"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	//Get the parameter
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")

	//Data verification
	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The phone num must be 11 digits!",
		})
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "Password cant less than 6 digits!",
		})
	}

	//If not input the name , we can give a 10-digit random string
	if len(name) == 0 {
		name = utils.RandomString(10)
	}

	log.Println(name, password, phone)
	ctx.JSON(utils.Succ("Register success!"))
}

func Login() {

}
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value // Get the value of the cookie
		service.DeleteSession(cookieValue)
		cookie.MaxAge = -1        // Set cookie invalidation
		http.SetCookie(w, cookie) // Send the modified cookie to the browser
	}
	GetPageBooksByPrice(w, r) // Go to homepage
}

// Regist Regist user
func Regist(w http.ResponseWriter, r *http.Request) {
	phone := r.PostFormValue("phone")
	password := r.PostFormValue("password")
	user, _ := service.CheckUserName(phone)
	if user.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "Username already exists")
	} else {
		service.SaveUser(phone, password)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "Incorrect username or password")
	}
}

// CheckUserName Verify that the username exists by sending an Ajax request
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	phone := r.PostFormValue("phone")
	user, _ := service.CheckUserName(phone)
	if user.ID > 0 {
		w.Write([]byte("already exist"))
	}
}
