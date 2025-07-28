package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestData struct {
	Valor int `json:"valor"`
}

type CoinbaseResponse struct {
	Data struct {
		Amount   string `json:"amount"`
		Base     string `json:"base"`
		Currency string `json:"currency"`
	}
}

func main() {
	r := gin.Default()

	r.GET("/converter", func(ctx *gin.Context) {
		var dados RequestData

		if err := ctx.ShouldBindJSON(&dados); err != nil {
			ctx.JSON(400, gin.H{"erro": err.Error()})
			return
		}

		resp, err := http.Get("https://api.coinbase.com/v2/prices/BTC-BRL/spot")
		if err != nil {
			ctx.JSON(500, gin.H{"erro": "Erro ao consultar API externa"})
			return
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			ctx.JSON(500, gin.H{"erro": "Erro ao ler resposta da API"})
			return
		}

		var cotacao CoinbaseResponse
		if err := json.Unmarshal(body, &cotacao); err != nil {
			ctx.JSON(500, gin.H{"erro": "Erro ao interpretar resposta da API"})
			return
		}

		precoBTC, err := strconv.ParseFloat(cotacao.Data.Amount, 64)
		if err != nil {
			ctx.JSON(500, gin.H{"erro": "Erro ao converter cotação"})
			return
		}

		valorBTC := float64(dados.Valor) / 100000000
		valorBRL := valorBTC * precoBTC

		ctx.JSON(200, gin.H{
			"cotacao":        valorBRL,
			"satoshis":       dados.Valor,
			"btc":            valorBTC,
			"brl_convertido": precoBTC,
		})

	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"menssagem": "Servidor funcionando!",
		})
	})

	r.Run(":8000")
}
