package main

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
)

func TestPingRoute(t *testing.T) {
    // 建立測試用的 Gin 路由引擎
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

    // 建立測試的 HTTP Request
    req, err := http.NewRequest(http.MethodGet, "/ping", nil)
    if err != nil {
        t.Fatal(err)
    }

    // 建立 Response 相應紀錄器
    res := httptest.NewRecorder()

    // 將 Request 發送到路由引擎進行處理
    r.ServeHTTP(res, req)

    // 檢查 Response 是否符合預期
    if res.Code != http.StatusOK {
        t.Errorf("expected status %d but got %d", http.StatusOK, res.Code)
    }
    expectedResponse := `{"message":"pong"}`
    if res.Body.String() != expectedResponse {
        t.Errorf("expected body %s but got %s", expectedResponse, res.Body.String())
    }
}
