package routes

import (
	"net/http"
	"strconv"

	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/database"
	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/models"
	"github.com/gin-gonic/gin"
)

func SetupArticlesRoutes(router *gin.Engine) {
    router.GET("/articles/:id", getArticle)
	router.GET("/articles/latests/:latests", getLatestArticles)
}

func getArticle(c *gin.Context) {
    id := c.Param("id")

    var article models.Article
    if err := database.DB.Where("id = ?", id).First(&article).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Text content not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"article": article})
}

func getLatestArticles(c *gin.Context) {
    latests, err := strconv.Atoi(c.Param("latests"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter 'latests'"})
        return
    }

    var articles []models.Article
    if err := database.DB.Order("id desc").Limit(latests).Find(&articles).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Articles not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"articles": articles})
}


