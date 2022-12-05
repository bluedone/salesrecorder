package service

import (
	"log"
	"net/http"
	DBEngine "sr-server/database"

	"github.com/gin-gonic/gin"
)

type SaleItem = DBEngine.SaleItem
type Sale = DBEngine.Sale

type saleItemSchema struct {
	ItemID int `json:"item_id"`
	Amount int `json:"amount"`
}

type addSaleSchema struct {
	Price float64          `json:"price"`
	Sales []saleItemSchema `json:"sales"`
}

func GetSaleItems(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func AddSale(c *gin.Context) {
	user_id := GetHeaderAuth(c)
	var addSale addSaleSchema
	if err := c.BindJSON(&addSale); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	db := DBEngine.CreateConnection()
	defer db.Close()
	var sale Sale
	sqlStatement := `INSERT INTO "sale" ("user_id", "created_at", "price") 	VALUES ($1, now(), $2)	RETURNING "id";`
	err := db.QueryRow(sqlStatement, user_id, addSale.Price).Scan(&sale.ID)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	sqlStatement = `INSERT INTO "sale_item" ("sale_id", "item_id", "amount") 	VALUES ($1, $2, $3) RETURNING "id";`
	var sale_item SaleItem
	for _, each := range addSale.Sales {
		err := db.QueryRow(sqlStatement, sale.ID, each.ItemID, each.Amount).Scan(&sale_item.ID)
		if err != nil {
			log.Fatalf("Unable to execute the query. %v", err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"id": sale.ID, "price": addSale.Price, "sales": addSale.Sales})
}
