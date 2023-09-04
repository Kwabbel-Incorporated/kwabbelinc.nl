package routes

import (
	"net/http"

	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/database"
	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/models"
	"github.com/gin-gonic/gin"
)

func SetupMediaRoutes(router *gin.Engine) {
    router.GET("/media/:id", getMedia)
}

func getMedia(c *gin.Context) {
    contentID := c.Param("id")

    var media models.Media
    if err := database.DB.Where("media_id = ?", contentID).First(&media).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Text content not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"media": media})
}
