Aqui está seu README atualizado, com melhorias para refletir melhor o que já foi implementado e algumas recomendações de padrões usados:

```markdown
# Order Food API - Golang + Gin + Design Patterns

Este projeto é um exemplo simples de uma API de pedidos de comida, desenvolvida em **Golang** com o framework **Gin**, aplicando padrões de projeto como **Factory** e **Repository**, além de boas práticas em arquitetura limpa.

---

## 🎯 Objetivo

Demonstrar na prática como aplicar o padrão **Factory** em uma API REST usando Go, permitindo a criação de diferentes tipos de pedidos (`regular` e `express`) a partir de uma única interface de criação, além de usar abstrações para persistência e manipulação dos dados.

---

## 📦 Funcionalidades

- Criar pedidos do tipo `regular` ou `express`
- Aplicar o padrão Factory para instanciar o tipo correto de pedido
- Organização modular do código com separação clara de responsabilidades (domínio, handler, repositório)
- Uso do padrão Repository para abstração da persistência (MongoDB)
- Tratamento básico de erros e validações

---

## 🗂️ Estrutura do Projeto

```
internal/
├── domain/         # Entidades de domínio, interfaces e lógica da Factory
├── handler/        # Handlers HTTP (controllers)
│   └── dto/        # Structs para requests e responses (DTOs)
├── repository/     # Implementações de persistência (MongoDB)
└── router/         # Definição das rotas e inicialização do servidor
cmd/
├── foodorder/main.go             # Ponto de entrada da aplicação
```

---

## 🚀 Como rodar

```bash
go mod tidy
go run cmd/foodorder/main.go
```

A API será executada em: [http://localhost:8080](http://localhost:8080)

---

## 🧪 Exemplos de Requisição cURL

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

- ✅ **Factory Pattern**: para criação dos pedidos `RegularOrder` e `ExpressOrder` via método construtor único
- ✅ **Repository Pattern**: abstração da persistência dos pedidos usando MongoDB
- ✅ **DTO (Data Transfer Object)**: para separar estrutura de dados entre camadas (HTTP e domínio)
- ✅ **Interface e Polimorfismo**: para manipular diferentes tipos de pedidos através da interface comum `Order`
- ✅ **Dependency Injection**: injeção do repositório no handler para facilitar testes e desacoplamento

---

## 📎 Requisitos

- Go 1.18 ou superior
- MongoDB rodando localmente ou remotamente (configurar URI no projeto)

---

## ✍️ Autor

Projeto criado para estudo e demonstração de padrões de projeto e boas práticas em desenvolvimento Go.

```

Se quiser, posso ajudar a gerar README com instruções para Docker, testes, ou outras ferramentas que você esteja usando. Quer?