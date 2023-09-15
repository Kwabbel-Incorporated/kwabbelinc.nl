package routes

import (
	"net/http"

	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/database"
	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/models"
	"github.com/gin-gonic/gin"
)

func SetupHeroRoutes(router *gin.Engine) {
    router.GET("/hero", getHero)
}

func getHero(c *gin.Context) {
    var hero models.Hero
    if err := database.DB.Where("is_current_hero = ?", 1).First(&hero).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Hero content not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": hero})
}

