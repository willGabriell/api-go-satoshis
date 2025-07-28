package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/willGabriell/api-go-satoshis/internal/model"
)

type CoinbaseClient struct{}

func NewCoinbaseClient() *CoinbaseClient {
	return &CoinbaseClient{}
}

func (c *CoinbaseClient) GetBTCPrice() (float64, error) {
	resp, err := http.Get("https://api.coinbase.com/v2/prices/BTC-BRL/spot")
	if err != nil {
		return 0, fmt.Errorf("erro ao consultar API externa: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("erro ao ler resposta da API: %w", err)
	}

	var cotacao model.CoinbaseResponse
	if err := json.Unmarshal(body, &cotacao); err != nil {
		return 0, fmt.Errorf("erro ao interpretar resposta da API: %w", err)
	}

	price, err := strconv.ParseFloat(cotacao.Data.Amount, 64)
	if err != nil {
		return 0, fmt.Errorf("erro ao converter cotação: %w", err)
	}

	return price, nil
}
