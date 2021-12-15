package router

import (
	"goauthtest/server2/controller"

	"github.com/gin-gonic/gin"
)

func OauthRouterInit(r *gin.Engine) {
	oauthRouter := r.Group("/oauth")
	{
		oauthRouter.GET("/authorize", controller.AuthorizeController)
		oauthRouter.POST("/token", controller.TokenController)
	}
}
func LoginRouterInit(r *gin.Engine) {
	loginRouter := r.Group("/")
	{
		loginRouter.GET("/login", controller.LoginHtmlController)
		loginRouter.POST("/login", controller.LoginController)
		loginRouter.GET("/auth", controller.AuthController)
	}
}
