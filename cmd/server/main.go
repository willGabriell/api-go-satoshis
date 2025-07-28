package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/willGabriell/api-go-satoshis/internal/handler"
)

func main() {
	r := gin.Default()

	// Configuração CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:3000", "http://127.0.0.1:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	convertHandler := handler.NewConvertHandler()

	r.POST("/converter", convertHandler.HandleConverter)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"menssagem": "Servidor funcionando!",
		})
	})

	r.Run(":8000")
}
