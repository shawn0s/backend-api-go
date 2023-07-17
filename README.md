
# Golang backend and Gin framework


```GO
package main

import (
	"fmt"
)

func main() {
	var a bool = true // Boolean

	b := 100 // Integer

	var (
		c float32 = 3.14  // Floating point number
		d string  = "Hi!" // String
	)

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)
	fmt.Println("map:", m)
}

}

```

## 建立GO backend project
```
$ mkdir ~/backend-api-go
$ go mod init backend-api-go
```

## 安裝gin 套件

添加 github.com/gin-gonic/gin 套件到專案
```
$ go get u github.com/gin-gonic/gin
```

## Running Gin 

```go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

```

添加Backend api

```GO
	v1 := r.Group("/api/v1")
	{
		v1.GET("/albums", album.GetAlbums)
	}
	
```




## 後端 啟動

```
# run main.go and visit 0.0.0.0:8080/ping on browser
$ go run main.go
```

## 新增 album RestFul api
想要建立一個專輯的功能，功能包含 新增 刪除 查詢 刪除，並添加至gin routing中




```GO
	v1 := router.Group("/api/v1")
	{

		v1.GET("/albums", album.GetAlbums)
	}


```




## Gin Middleware

### Gin with Auth Middleware  
添加


```GO
func ValidateToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	log.Println("AuthMiddleware" + token)
	if token == "" {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}
	
	// 建立 HTTP 客戶端
	client := &http.Client{}

	// 建立 GET 請求
	req, err := http.NewRequest("POST", "http://localhost:8080", nil)
	req.Header.Add("Authorization", token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}

	// 發送請求
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}
	c.Set("userAuthDTO", resp)
}

```
### Gin with CORS Middleware
添加CORS CORS Settings 包裝成middleware 讓 gin.Default() 使用

```GO
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
```


## 新增Swagger 服務

## go-stress-test

```sh
curl --location 'http://localhost:8088/api/v1/albums' \
--header 'accept: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzUxMiJ9.eyJjcmVhdGVkIjoxNjg5NDk5NjE5NTMxLCJleHAiOjE2OTIwOTE2MTksInVzZXJJZCI6MSwiZW1haWwiOiJhY2FkZW15LmRldkBkZWx0YXd3LmNvbSJ9.-EasziixGocxOPzqYahhQpOuBUgZI6UpyOo0Ih3PDwCiapiFjAezwSGfRy7umQt2hJyicqZPcfz-i6sHe0X1fA'
```