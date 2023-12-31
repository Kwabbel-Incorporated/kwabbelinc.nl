package routes

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/database"
	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/models"
	"github.com/gin-gonic/gin"
)

func SetupArticlesRoutes(router *gin.Engine) {
    router.GET("/articles/:id", getArticle)
	router.GET("/articles/latests/:latests", getLatestArticles)
    router.GET("/articles/content/:id", getArticleContent)
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

func getArticleContent(c *gin.Context) {
    id := c.Param("id")

    filePath := filepath.Join("public", "articles", id+".html")

    content, err := os.ReadFile(filePath)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "HTML-bestand niet gevonden"})
        return
    }

    c.Data(http.StatusOK, "text/html", content)
}


