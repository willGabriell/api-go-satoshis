package main

import (
	"github.com/gin-gonic/gin"
	"github.com/willGabriell/api-go-satoshis/internal/handler"
)

func main() {
	r := gin.Default()

	convertHandler := handler.NewConvertHandler()

	r.GET("/converter", convertHandler.HandleConverter)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"menssagem": "Servidor funcionando!",
		})
	})

	r.Run(":8000")
}
