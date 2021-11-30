package router

import (
	"zimu-shop/controller"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("api/user/register", controller.Register)
	r.Run()

}
