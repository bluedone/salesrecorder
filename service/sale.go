package service

import (
	"net/http"
	DBEngine "sr-server/database"

	"github.com/gin-gonic/gin"
)

type SaleItem = DBEngine.SaleItem
type Sale = DBEngine.Sale

func GetSaleItems(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
