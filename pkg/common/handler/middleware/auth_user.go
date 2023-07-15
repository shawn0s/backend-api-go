package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func auth(token string) {

}

// 自定義驗證 Middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在這裡進行使用者驗證的邏輯，例如檢查驗證 token

		// 假設驗證失敗，返回 401 錯誤
		if !authenticateUser(c.Request) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
			c.Abort() // 中止後續處理函式的執行
			return
		}

		// 驗證成功，繼續執行下一個處理函式
		c.Next()
	}
}

// 處理 /hello 路由的函式
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, authenticated user!"})
}
