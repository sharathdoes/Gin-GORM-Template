package initializers

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
) 
var DB *gorm.DB
func ConnectWithDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ENV File is Missing")
	}
	dsn:=os.Getenv("DATABASE_URL")
	if dsn=="" {
		log.Fatal("DATABASE_URL Not Found")
	}
	var err2 error
	DB, err2 = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err2 != nil {
		log.Fatal(" Failed to connect to Supabase:", err2)
	}

	fmt.Println(" Connected to  PostgreSQL successfully!",DB)
} 