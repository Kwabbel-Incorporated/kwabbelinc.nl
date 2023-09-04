package server

import (
	"fmt"
	"log"
	"os"

	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/server/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	// gin.SetMode(gin.ReleaseMode)

    router := gin.Default()

    config := cors.DefaultConfig()
    config.AllowOrigins = []string{fmt.Sprintf("http://%s:%s", os.Getenv("CORS_API_HOST"), os.Getenv("CORS_API_PORT"))}
    config.AllowHeaders = []string{"Authorization"}

    router.Use(cors.New(config))

    routes.SetupTextContentRoutes(router)
	routes.SetupArticlesRoutes(router)
    routes.SetupMediaRoutes(router)

    err := router.Run("localhost:8080")
    if err != nil {
        log.Fatal("Failed to start server:", err)
    }
}