package main

import (
	"backend-api-go/pkg/album"

	docs "backend-api-go/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	httpPort = "8088"
)



func main() {
	// market.TradingFund()
	// docs.SwaggerInfo.BasePath = "/api/v1"
	// user.CreateUSer()
	router := gin.Default()

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:"+httpPort
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.Use()
	// router.GET("/test", user.Test)
	router.GET("/albums", album.GetAlbums)






	url := ginSwagger.URL("http://localhost:"+ httpPort +"/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
	router.Run(":" + httpPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
