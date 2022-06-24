package postgresDB

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
	"os"
)

// Init setting up mongDB
func Init() *gorm.DB {
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")
	DBHost := os.Getenv("DB_HOST")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")
	DBTimeZone := os.Getenv("DB_TIMEZONE")
	DBMode := os.Getenv("DB_MODE")
	var dsn string
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", DBHost, DBUser, DBPass, DBName, DBPort, DBMode, DBTimeZone)
	} else {
		dsn = databaseUrl
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	log.Println("Database Connected Successfully...")
	return db
}
