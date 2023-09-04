package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/models"
	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/util"
)

var (
	db *gorm.DB
	err error
)

func main() {
	util.LoadEnvVariables()
	connectDatabase()
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
  
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowHeaders = []string{"Authorization"}
  
	router.Use(cors.New(config))
  
	router.GET("/textcontent/:id", getTextContent)
  
	err := router.Run("localhost:8080")
	if err != nil {
	  log.Fatal("Failed to start server:", err)
	}
  }

func connectDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.TextContent{})
	log.Println("Succeeded in connecting to the database")
}

func getTextContent(c *gin.Context) {
    contentID := c.Param("id")

    var textContent models.TextContent // Gebruik models.TextContent
    if err := db.Where("content_id = ?", contentID).First(&textContent).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Text content not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"content_text": textContent.ContentText})
}
