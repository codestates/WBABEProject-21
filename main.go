package main

import (
	"pkg/config"
	conf "pkg/config"
	"pkg/router"
	ut "pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.SetupRouter(r)
	// # DB Connection & Router
	config.ConnectDB()
	// # Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(router.CORS())
	// # Configs
	toml := conf.GetConfig("./config/config.toml")
	err := r.Run(toml.Server.Port)
	ut.HandleErr(err)
}
