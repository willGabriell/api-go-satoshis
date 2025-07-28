# API Go Satoshis

Uma API REST simples em Go para conversão de Satoshis para Real Brasileiro (BRL), utilizando a cotação atual do Bitcoin da Coinbase.

## Funcionalidades

- Conversão de Satoshis para BRL em tempo real
- Integração com API da Coinbase para cotação do Bitcoin
- API REST com endpoint simples
- Containerização com Docker

## Tecnologias

- **Go 1.23** - Linguagem de programação
- **Gin** - Framework web HTTP
- **Docker** - Containerização

## Estrutura do Projeto

```
api-go-satoshis/
├── cmd/server/          # Ponto de entrada da aplicação
├── internal/
│   ├── client/         # Cliente para APIs externas (Coinbase)
│   ├── handler/        # Handlers HTTP
│   ├── model/          # Estruturas de dados
│   └── service/        # Lógica de negócio
├── Dockerfile          # Configuração Docker
└── README.md
```

## Como executar

### Localmente

```bash
# Clone o repositório
git clone https://github.com/willGabriell/api-go-satoshis.git
cd api-go-satoshis

# Execute a aplicação
go run cmd/server/main.go
```

### Com Docker

```bash
# Build da imagem
docker build -t api-go-satoshis .

# Execute o container
docker run -p 8000:8000 api-go-satoshis
```

A API estará disponível em `http://localhost:8000`

## Endpoints

### GET /
Endpoint de health check

**Resposta:**
```json
{
  "menssagem": "Servidor funcionando!"
}
```

### GET /converter
Converte Satoshis para Real Brasileiro

**Exemplo de requisição:**
```json
{
  "valor": 100000000
}
```

**Exemplo de resposta:**
```json
{
  "cotacao": 350000.50,
  "satoshis": 100000000,
  "btc": 1.0,
  "brl_convertido": 350000.50
}
```

## Como funciona

1. A API recebe uma requisição com a quantidade de Satoshis
2. Consulta a cotação atual do Bitcoin na API da Coinbase
3. Converte Satoshis para BTC (1 BTC = 100.000.000 Satoshis)
4. Multiplica o valor em BTC pela cotação em BRL
5. Retorna o resultado da conversão

## Desenvolvimento

O projeto segue uma arquitetura limpa com separação de responsabilidades:

- **Handler**: Gerencia requisições HTTP
- **Service**: Contém a lógica de negócio
- **Client**: Integração com APIs externas
- **Model**: Estruturas de dados

## Autor

[@willGabriell](https://github.com/willGabriell)