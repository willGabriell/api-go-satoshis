package model

type RequestData struct {
	Valor int `json:"valor"`
}

type ConvertResponse struct {
	Cotacao       float64 `json:"cotacao"`
	Satoshis      int     `json:"satoshis"`
	BTC           float64 `json:"btc"`
	BRLConvertido float64 `json:"brl_convertido"`
}
