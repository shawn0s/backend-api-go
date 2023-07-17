package main

import (
	docs "backend-api-go/docs"
	"backend-api-go/pkg/album"
	"backend-api-go/pkg/common/handler/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	httpPort = "8088"
	DbUrl    = "postgresql://postgres:123456@backend-postgresql-headless:5432/dkam-order"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	router := gin.New()

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + httpPort
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	v1 := router.Group("/api/v1")
	{
		v1.Use(middleware.ValidateToken)
		v1.Use(middleware.Logger())
		v1.GET("/albums", album.GetAlbums)
	}

	url := ginSwagger.URL("http://localhost:" + httpPort + "/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	router.Run(":" + httpPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
