package initializers

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func ConnectWithDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ENV File is Missing")
	}
	dsn:=os.Getenv("DATABASE_URL")
	if dsn=="" {
		log.Fatal("DATABASE_URL Not Found")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(" Failed to connect to Supabase:", err)
	}

	fmt.Println(" Connected to Supabase PostgreSQL successfully!",db)
} 