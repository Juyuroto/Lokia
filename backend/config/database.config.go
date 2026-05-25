package config

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/backend/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := os.Getenv("PostgreHost")
    port := os.Getenv("PostgrePort")
    user := os.Getenv("PostgreUser")
    password := os.Getenv("PostgrePassword")
    dbname := os.Getenv("PostgreDB")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
                    	host, user, password, dbname, port,)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
    	log.Fatal("Impossible de se connecter à la base de données :", err)
    }

    fmt.Println("Connexion à la base de données réussie !")
    
    db.AutoMigrate(&models.History{}, &models.Conversation{})

    DB = db
}