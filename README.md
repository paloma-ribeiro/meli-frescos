# Meli Frescos

<p align="center">
  <a href="https://github.com/paloma-ribeiro/meli-frescos/actions/workflows/test.yml">
    <img src="https://github.com/paloma-ribeiro/meli-frescos/actions/workflows/test.yml/badge.svg">
  </a>
</p>

### Inicialização da aplicação

> go mod init github.com/paloma-ribeiro/meli-frescos

### Atualização das importações

> go mod tidy

### Instalação do drive sql

> go get -u github.com/go-sql-driver/mysql

### Instalação do Gin

> go get -u github.com/gin-gonic/gin

### Instalação GoDotEnv

> go get github.com/joho/godotenv

### Instalação do Swagger

> go get -u github.com/swaggo/swag/cmd/swag

> go get -u github.com/swaggo/files

> go get -u github.com/swaggo/gin-swagger

#### Gerar documentação com swagger

> swag init -g cmd/server/main.go