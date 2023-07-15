
# Golang backend and Gin framework


## 建立GO專案
```
$ mkdir ~/backend-api-go
$ go mod init backend-api-go
```

## 安裝gin 套件
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

## 後端啟動

```
# run main.go and visit 0.0.0.0:8080/ping on browser
$ go run main.go
```

## 創建al

## 套用 postgresql db


## 新增Swagger 服務

## Middleware

### Cors

### Auth

### Logger


## Swagger

安裝swag cmd
```
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
```


```
swag init
```

## go-stress-test