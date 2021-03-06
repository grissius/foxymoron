package api

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	_ "github.com/grissius/foxymoron/api"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Foxymoron REST API
// @version 1.0
// @description API Proxy to GitLab

// @license.name MIT

// @host foxymoron.appspot.com
// @BasePath /

// @securityDefinitions.apikey ApiKey
// @in header
// @name Authorization

// @securityDefinitions.apikey GitLabURL
// @in header
// @name X-Gitlab-Url
func createEngine() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AddAllowHeaders("*")
	r.Use(cors.New(config))
	r.GET("/", root)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	proxy := r.Group("/")
	proxy.Use(authMdw)
	{
		proxy.GET("/projects", getProjectsController)
		proxy.GET("/commits", getCommitsController)
		proxy.GET("/statistics", getStatisticsController)
	}
	return r
}

func RunAt(port int) {
	log.Printf("Startig server on port %v", port)
	createEngine().Run(":" + strconv.Itoa(port))
}
