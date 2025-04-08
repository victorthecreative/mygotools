# Como usar o NewConfig

O módulo de configuração permite carregar configurações a partir de diferentes formatos de arquivo (`.env`, `.yaml` ou `.json`).

## Instalação

```bash
go get github.com/spf13/viper
go get github.com/cockroachdb/errors
```

## Uso Básico

```go
package main

import (
	"fmt"
	"log"

	"seu/pacote/config"
)

func main() {
	// Carrega configurações a partir de um arquivo .env
	cfg, err := config.NewConfig("env")
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	fmt.Printf("Porta do serviço: %s\n", cfg.Port)
	fmt.Printf("URL do serviço: %s\n", cfg.ServiceUrl)

	// Para conexão com o Postgres
	fmt.Printf("Host do Postgres: %s\n", cfg.PostgresHost)

	// Para conexão com o MongoDB
	fmt.Printf("URI do MongoDB: %s\n", cfg.MongoURI)

	// Para conexão com o Redis
	fmt.Printf("Host do Redis: %s\n", cfg.RedisHost)
}
```

## Formatos Suportados

O parâmetro `typeConfig` define qual tipo de arquivo de configuração será usado:

- `"env"`: Carrega de um arquivo `.env` no diretório atual
- `"yaml"`: Carrega de um arquivo `config_yaml.yaml` no diretório atual
- `"json"`: Carrega de um arquivo `config_json.json` no diretório atual

## Estrutura dos Arquivos de Configuração

### Arquivo .env

```
PORT=8080
SERVICE_URL=localhost:8080
POSTGRES_HOST=localhost
MONGO_URI=mongodb://localhost:27017
REDIS_HOST=localhost
```

### Arquivo config_yaml.yaml

```yaml
Port: "8080"
ServiceUrl: "localhost:8080"
PostgresHost: "localhost"
MongoURI: "mongodb://localhost:27017"
RedisHost: "localhost"
```

### Arquivo config_json.json

```json
{
  "Port": "8080",
  "ServiceUrl": "localhost:8080",
  "PostgresHost": "localhost",
  "MongoURI": "mongodb://localhost:27017",
  "RedisHost": "localhost"
}
```

## Nota Importante

Certifique-se de que o arquivo de configuração escolhido existe no diretório atual, pois a função `NewConfig` retornará um erro caso não consiga encontrar o arquivo especificado.

As variáveis de ambiente têm precedência sobre as configurações definidas nos arquivos. Isso significa que, se uma variável de ambiente estiver definida, ela sobrescreverá o valor correspondente no arquivo de configuração.
