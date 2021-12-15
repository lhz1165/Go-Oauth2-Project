package main

import (
	"goauthtest/server2/config"
	rr "goauthtest/server2/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitServer()

	router := gin.Default()
	router.LoadHTMLGlob("static/*")
	router.Static("/static", "./static")
	rr.OauthRouterInit(router)
	rr.LoginRouterInit(router)

	router.Run(":9096")
}
