package middleware

import (
	"io"
	"net/http"
	"encoding/json"
	"log"
	"github.com/gin-gonic/gin"
)

type UserAuthDTO struct {
	Id int 				`json:"id"`
	Name string 		`json:"name"`
	Email string 		`json:"email"`
	Phone string 		`json:"phone"`
	Title string 		`json:"title"`
	UserType string 	`json:"userType"`
	Avatar string 		`json:"avatar"`
	Description string 	`json:"description"`
	Tags  string  		`json:"tags"`

}

// ValidateToken returns the token's claims if the token is valid
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

	// values := map[string]string{}
	// jsonValue, _ := json.Marshal(values)
	// 建立 GET 請求
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/bfb/auth", nil)
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
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
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var result UserAuthDTO
	json.Unmarshal(body, &result)
	
	c.Set("userAuthDTO", result)
	
}
