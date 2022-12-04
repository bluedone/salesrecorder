package service

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"

	DBEngine "sr-server/database"

	"github.com/gin-gonic/gin"
)

type User = DBEngine.User
type Token = DBEngine.AccessToken

func CreateUser(c *gin.Context) {
	AdminAuth(c)
	var user User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	hashPassword := sha256.New()
	hashPassword.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hashPassword.Sum((nil)))

	db := DBEngine.CreateConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO "user" ("username", "password", "created_at") VALUES ($1, $2, now()) RETURNING "id", "username", "created_at"`
	err := db.QueryRow(sqlStatement, user.Username, user.Password).Scan(&user.ID, &user.Username, &user.CreatedAt)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	var access Token
	sqlStatement = `INSERT INTO "access_token" ("user_id", "token") VALUES ($1, $2) RETURNING "id", "user_id", "token"`
	err = db.QueryRow(sqlStatement, user.ID, "").Scan(&access.ID, &access.UserID, &access.Token)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	log.Printf("Inserted a single record %v", user.ID)
	c.JSON(http.StatusCreated, gin.H{"username": user.Username, "id": user.ID, "created_at": user.CreatedAt})
}

func LogIn(c *gin.Context) {
	var requestedUser User
	var user User
	if err := c.BindJSON(&requestedUser); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	db := DBEngine.CreateConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM "user" WHERE "username" = ($1)`

	err := db.QueryRow(sqlStatement, requestedUser.Username).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	hashPassword := sha256.New()
	hashPassword.Write([]byte(requestedUser.Password))
	requestedUser.Password = hex.EncodeToString(hashPassword.Sum((nil)))
	if requestedUser.Password != user.Password {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Wrong username or password"})
		return
	} else {
		token := GenerateSecureToken(db, user.ID)
		c.JSON(http.StatusOK, gin.H{"id": user.ID, "username": user.Username, "created_at": user.CreatedAt, "token": token})
	}
}

func GenerateSecureToken(db *sql.DB, user_id int) string {
	b := make([]byte, 100)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	token := hex.EncodeToString(b)

	sqlStatement := `UPDATE "access_token" SET	"id" = "id", "user_id" = $1, "token" = $2 WHERE "user_id" = $1;`
	_, err := db.Exec(sqlStatement, user_id, token)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	return token
}

func GetUserFromToken(token string) int {
	var access Token

	db := DBEngine.CreateConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM "access_token" WHERE "token" = ($1)`

	err := db.QueryRow(sqlStatement, token).Scan(&access.ID, &access.UserID, &access.Token)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	log.Printf("Authorized User %v", access.UserID)
	return access.UserID
}
