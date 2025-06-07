# Order Food API - Golang + Gin + Design Patterns

Este projeto é um exemplo simples de uma API de pedidos de comida, desenvolvida em **Golang** com o framework **Gin**, aplicando o padrão de projeto **Factory**.

---

## 🎯 Objetivo

Demonstrar na prática como aplicar o padrão **Factory** em uma API REST usando Go, permitindo a criação de diferentes tipos de pedidos (`regular` e `express`) a partir de uma única interface de criação.

---

## 📦 Funcionalidades

- Criar pedidos do tipo `regular` ou `express`
- Aplicar o padrão Factory para instanciar o tipo correto de pedido
- Organização modular do código com separação por responsabilidade

---

## 🗂️ Estrutura do Projeto

```
internal/
├── domain/         # Entidades de domínio e a lógica da Factory
├── handler/        # Handlers HTTP (controllers)
│   └── dto/        # Structs de request e response
└── router/         # Definição das rotas
cmd/
├── foodorder/main.go             # Ponto de entrada da aplicação
```

---

## 🚀 Como rodar

```bash
go mod tidy
go run main.go
```

A API será executada em: [http://localhost:8080](http://localhost:8080)

---

## 🧪 Exemplo de Requisição cURL

### Criar pedido regular

```bash
curl -X POST http://localhost:8080/v1/orders \
  -H "Content-Type: application/json" \
  -d '{
    "order_type": "regular",
    "customer_id": "123",
    "items": [
      {
        "product_id": "p1",
        "name": "Pizza",
        "price": 30.0,
        "quantity": 1
      }
    ]
  }'
```

### Criar pedido express

```bash
curl -X POST http://localhost:8080/v1/orders \
  -H "Content-Type: application/json" \
  -d '{
    "order_type": "express",
    "customer_id": "123",
    "items": [
      {
        "product_id": "p1",
        "name": "Pizza",
        "price": 30.0,
        "quantity": 1
      }
    ],
    "delivery_limit": 1800
  }'
```

---

## 📌 Padrões Aplicados

- ✅ **Factory Pattern**: responsável por criar instâncias de `Order` conforme o tipo (`regular` ou `express`)

---

## 📎 Requisitos

- Go 1.18 ou superior

---

## ✍️ Autor

Desenvolvido para fins de estudo e demonstração de boas práticas em Go.