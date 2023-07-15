package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUSer() {

}

func Test(c *gin.Context) {
	str := []byte("ok")                      // 對於[]byte感到疑惑嗎？ 因為網頁傳輸沒有string的概念，都是要轉成byte字節方式進行傳輸
	c.Data(http.StatusOK, "text/plain", str) // 指定contentType為 text/plain，就是傳輸格式為純文字啦～
}

func CreateUser(c *gin.Context) {

}

func QueryUsers(c *gin.Context) {

	// query
	queryStmt := `update "Students" set "Name"=$1, "Roll_Number"=$2 where "id"=$3`
	_, e := db.Exec(updateStmt, "Rachel", 24, 8)
	CheckError(e)

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
}

func UpdateUser(c *gin.Context) {
	// update
	updateStmt := `update "Students" set "Name"=$1, "Roll_Number"=$2 where "id"=$3`
	_, e := db.Exec(updateStmt, "Rachel", 24, 8)
	CheckError(e)

	str := []byte("ok")
	c.Data(http.StatusOK, "text/plain", str)
}

func DeleteUser(c *gin.Context) {

}
