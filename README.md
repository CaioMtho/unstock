# Uncomplicated Stock

Uma aplicação simples de controle de estoque com suporte a alertas em tempo real via WebSocket. Desenvolvida em Go, com uma interface web leve e interativa.

---

## Funcionalidades

- API REST para gerenciamento de produtos
- Atualização de estoque com validação de mínimo
- Alertas em tempo real via WebSocket
- Interface web básica
- Binários multiplataforma gerados via Go

---

## Requisitos

> Para build e execução:

- [Go](https://golang.org/dl/) (versão 1.25.3 ou superior)

---

## Como usar

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/unstock.git
cd unstock
```

### 2. Instale as dependências (se houver)

```bash
go mod tidy
```

### 3. (Opcional) Compile os binários

```bash
GOOS=linux GOARCH=amd64 go build -o unstock-linux ./cmd/unstock
GOOS=windows GOARCH=amd64 go build -o unstock.exe ./cmd/unstock
GOOS=darwin GOARCH=arm64 go build -o unstock-mac ./cmd/unstock
```

### 4. Execute

```bash
# caso esteja na raiz do projeto
go run ./cmd
```

Acesse a interface no navegador:

```bash
http://localhost:8080
```

## Endpoints Principais

- GET /api/products/ - Lista todos os produtos
- POST /api/products/ - Cria um novo produto
- PATCH /api/products/:id/stock - Atualiza o estoque de um produto
- DELETE /api/products/:id - Remove um produto
- GET /ws/alerts - WebSocket para alertas de estoque
