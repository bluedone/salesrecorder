package service

import (
	"log"
	"net/http"

	DBEngine "sr-server/database"

	"github.com/gin-gonic/gin"
)

type Item = DBEngine.Item
type AccessToken = DBEngine.AccessToken

func GetItems(c *gin.Context) {
	AdminAuth(c)
	db := DBEngine.CreateConnection()
	defer db.Close()
	var items []Item
	sqlStatement := `SELECT * FROM item`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var item Item
		err = rows.Scan(&item.ID, &item.Name, &item.CreatedAt, &item.Cost, &item.UserID)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		items = append(items, item)
	}
	c.JSON(http.StatusOK, items)
}

func GetItemsByUser(c *gin.Context) {
	user_id := GetHeaderAuth(c)
	db := DBEngine.CreateConnection()
	defer db.Close()
	var items []Item
	sqlStatement := `SELECT * FROM item WHERE user_id = ($1) `
	rows, err := db.Query(sqlStatement, user_id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var item Item
		err = rows.Scan(&item.ID, &item.Name, &item.CreatedAt, &item.Cost, &item.UserID)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		items = append(items, item)
	}
	c.JSON(http.StatusOK, items)
}

func CreateItem(c *gin.Context) {
	AdminAuth(c)
	var item Item

	if err := c.BindJSON(&item); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := DBEngine.CreateConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO "item" ("name", "created_at", "cost", "user_id") 	VALUES ($1, now(), $2, $3)	RETURNING "id", "name", "created_at", "cost", "user_id";`
	err := db.QueryRow(sqlStatement, item.Name, item.Cost, item.UserID).Scan(&item.ID, &item.Name, &item.CreatedAt, &item.Cost, &item.UserID)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	log.Printf("Inserted a single record %v", item.ID)
	c.JSON(http.StatusCreated, item)
}
