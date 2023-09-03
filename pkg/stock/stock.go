package stock

import (
	"backend-api-go/pkg/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

type StockRequestBody struct {
	StockNo            string    `json:"stockNo"`
	StockName          string    `json:"stockName"`
	StockType          string    `json:"stockType"`
	StockDate          string    `json:"stockDate"`
	Date               time.Time `json:"date"`
	Amount             int       `json:"amount"`
	Avg                float32   `json:"avg"`
	Change             float32   `json:"change"`
	Close              float32   `json:"close"`
	Diff               float32   `json:"diff"`
	High               float32   `json:"high"`
	Low                float32   `json:"low"`
	Open               float32   `json:"open"`
	PeRatio            float32   `json:"peRatio"`
	Volume             int       `json:"volume"`
	DealerBuy          int       `json:"dealerBuy"`
	DealerSell         int       `json:"dealerSell"`
	DealerTotal        int       `json:"dealerTotal"`
	DealerTotalAmount  int       `json:"dealerTotalAmount"`
	ForeignBuy         int       `json:"foreignBuy"`
	ForeignSell        int       `json:"foreignSell"`
	ForeignTotal       int       `json:"foreignTotal"`
	ForeignTotalAmount int       `json:"foreignTotalAmount"`
	FundBuy            int       `json:"fundBuy"`
	FundSell           int       `json:"fundSell"`
	FundTotal          int       `json:"fundTotal"`
	FundTotalAmount    int       `json:"fundTotalAmount"`
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/stock")
	routes.POST("/", h.CreateStock)
	routes.GET("/:id", h.Getstock)
	routes.PUT("/:id", h.UpdateStock)
	routes.DELETE("/:id", h.DeleteStock)
	routes.GET("/:id", h.QueryStock)
}

func (h handler) CreateStock(ctx *gin.Context) {
	body := StockRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var stock models.Stock

	stock.StockNo = body.StockNo

	if result := h.DB.Create(&stock); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &stock)
}

func (h handler) Getstock(ctx *gin.Context) {

}

func (h handler) UpdateStock(ctx *gin.Context) {

}

func (h handler) DeleteStock(ctx *gin.Context) {

}

func (h handler) QueryStock(ctx *gin.Context) {

}
