package routes

import (
	"net/http"

	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/database"
	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/models"
	"github.com/gin-gonic/gin"
)

func SetupTextContentRoutes(router *gin.Engine) {
    router.GET("/textcontent/:id", getTextContent)
}

func getTextContent(c *gin.Context) {
    contentID := c.Param("id")

    var textContent models.TextContent
    if err := database.DB.Where("content_id = ?", contentID).First(&textContent).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Text content not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"content_text": textContent.ContentText})
}
