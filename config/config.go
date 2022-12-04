package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	Err                = godotenv.Load("local.env")
	Origins            = strings.Split(os.Getenv("CORS"), ",")
	postgresDB         = os.Getenv("POSTGRES_DB")
	postgresPASSWORD   = os.Getenv("POSTGRES_PASSWORD")
	postgresUSER       = os.Getenv("POSTGRES_USER")
	postgresHOST       = os.Getenv("POSTGRES_HOST")
	postgresPORT, _err = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	DatabaseURL        = fmt.Sprintf("postgresql://%v:%v@%v:%v/%v", postgresUSER, postgresPASSWORD, postgresHOST, postgresPORT, postgresDB)
	PsqlInfo           = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", postgresHOST, postgresPORT, postgresUSER, postgresPASSWORD, postgresDB)
	AdminUser          = os.Getenv("ADMIN_USER")
	AdminPassword      = os.Getenv("ADMIN_PASSWORD")
)
