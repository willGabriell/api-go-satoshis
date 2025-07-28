package service

import (
	"fmt"

	"github.com/willGabriell/api-go-satoshis/internal/client"
	"github.com/willGabriell/api-go-satoshis/internal/model"
)

type ConverterService struct {
	coinbaseClient *client.CoinbaseClient
}

func NewConverterService() *ConverterService {
	return &ConverterService{
		coinbaseClient: client.NewCoinbaseClient(),
	}
}

func (s *ConverterService) ConvertSatoshisToBRL(satoshis int) (*model.ConvertResponse, error) {
	precoBTC, err := s.coinbaseClient.GetBTCPrice()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter cotação do BTC: %w", err)
	}

	valorBTC := float64(satoshis) / 100000000

	valorBRL := valorBTC * precoBTC

	return &model.ConvertResponse{
		Cotacao:       precoBTC,
		Satoshis:      satoshis,
		BTC:           valorBTC,
		BRLConvertido: valorBRL,
	}, nil
}
