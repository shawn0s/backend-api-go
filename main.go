package main

import (
	docs "backend-api-go/docs"
	"backend-api-go/pkg/common/db"
	"backend-api-go/pkg/stock"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
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

	router := gin.Default()

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + httpPort
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	dbHandler := db.InitDB(os.Getenv("POSTGRESQL_URL"))

	stock.RegisterRoutes(router, dbHandler)

	url := ginSwagger.URL("http://localhost:" + httpPort + "/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	router.Run(":" + os.Getenv("SEVER_PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
