<div align="center">
    <img alt="Golang Logo" title="Golang Logo" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" width="300" height="300">
</div>

# BOILERPLATE GOLANG - AUTH

## O que é necessário para testar essa aplicação?

- Ter o [Golang](https://go.dev/) instalado;
- Algum Rest Client => [Insomnia](https://insomnia.rest/download), [Postman](https://www.postman.com/), [cURL](https://curl.se/), [Thunder](https://www.thunderclient.com/);

## Libs utilizadas

- [GIN Framework HTTP](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [GORM Driver PostgreSQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)
- [JWT](https://pkg.go.dev/github.com/golang-jwt/jwt/v4)
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [GoDotEnv](https://pkg.go.dev/github.com/joho/godotenv)
- [Swaggo](https://github.com/swaggo/swag)

## Como instalar?

- GIN Framework HTTP =>

```bash
go get -u github.com/gin-gonic/gin
```

- GORM =>

```bash
go get -u gorm.io/gorm
```

- GORM Driver PostgreSQL =>

```bash
go get -u gorm.io/driver/postgres
```

- JWT =>

```bash
go get -u github.com/golang-jwt/jwt/v4
```

- Bcrypt =>

```bash
go get -u golang.org/x/crypto/bcrypt
```

- GoDotEnv =>

```bash
go get github.com/joho/godotenv
```

- Go Swagger =>

```bash
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

## Variáveis de Ambiente

Para rodar esse projeto, você vai precisar adicionar as seguintes variáveis de ambiente no seu .env

`PORT`

`DB`

`JWT_SECRET_KEY`

## Rodando localmente

Clone o projeto

```bash
  git clone https://link-para-o-projeto
```

Entre no diretório do projeto

```bash
  cd template-go-kukac
```

Instale as dependências

```bash
  go mid tidy
```

Inicie o servidor

```bash
  go run main.go
```

## Funcionalidades

- [x] Cadastro do usuário;
- [x] Login de um usuário cadastrado;
- [x] Armazena o token JWT do usuário logado nos cookies;
- [x] Valida se o usuário está realmente logado;

## Endpoints

#### Swagger Documentation

```http
  GET /docs/index.html
```

#### Cadastro

```http
  POST /signup
```

| Parâmetro | Tipo     | Descrição                         |
| :-------- | :------- | :-------------------------------- |
| `email`   | `string` | **Obrigatório**. Email do Usuário |
| `senha`   | `string` | **Obrigatório**. Senha do Usuário |

#### Login

```http
  POST /login
```

| Parâmetro | Tipo     | Descrição                         |
| :-------- | :------- | :-------------------------------- |
| `email`   | `string` | **Obrigatório**. Email do Usuário |
| `senha`   | `string` | **Obrigatório**. Senha do Usuário |

#### Validate

```http
  GET /validate
```

| Parâmetro | Tipo   | Descrição |
| :-------- | :----- | :-------- |
| `none`    | `none` | none      |

- Exemplo de Payload:

```json
{
  "email": "teste@email.com",
  "password": "senha123"
}
```

- Exemplo de Requisição:

```sh
curl -X POST \
  http://localhost:3000/signup \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{
	"email": "teste@email.com",
    "password":"senha123",
}'
```
