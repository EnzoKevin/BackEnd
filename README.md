# 🚀 BackEnd – API em Go

API REST desenvolvida em **Go** seguindo uma arquitetura em camadas (**Handlers → UseCases → Repositories**), com foco em organização, escalabilidade e boas práticas de backend.

Este projeto implementa operações básicas de **gerenciamento de usuários**, servindo como base para aplicações maiores (ex: autenticação, produtos, carrinho de compras, etc).



## 📌 Tecnologias Utilizadas

- **Go (Golang)**
- **net/http**
- **database/sql**
- **UUID (github.com/google/uuid)**
- Arquitetura em camadas (Clean Architecture simplificada)



## 📂 Estrutura do Projeto

BackEnd/

BackEnd/<br>
├── client/<br>
├── internal/<br>
│   ├── handlers/<br>
│   ├── usecases/<br>
│   ├── repositories/<br>
│   │   └── user/<br>
│   └── models/<br>
├── main.go<br>
├── go.mod<br>
└── README.md

## ⚙️ Pré-requisitos

- Go **1.18+**
- Banco de dados relacional (ex: PostgreSQL)
- Git



##  Como Executar o Projeto

###  Clone o repositório

```bash
git clone https://github.com/EnzoKevin/BackEnd.git
cd BackEnd

go mod tidy

```

### Configure o banco de dados

Certifique-se de que:

O banco está rodando

A string de conexão está correta no código de inicialização

### Execute a Aplicação 

```bash
  go run main.go
```


### Endpoints Disponíveis
🔹 Criar Usuário

POST /users

Request Body:
```json

{
  "name": "John Doe",
  "email": "john@doe.com",
  "password": "123456"
}
```


🔹 Buscar Usuário

GET /users

Request Body:
```json

{
  "id": 1
  "name": "John Doe",
  "email": "john@doe.com",
  "password": "123456"
}
```
