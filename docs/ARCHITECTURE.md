# 📊 Arquitetura Visual do Sistema FinTrack

## 🏗️ Visão Geral do Sistema

```
┌─────────────────────────────────────────────────────────────────┐
│                         USUÁRIO                                  │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│                    FRONTEND (Next.js)                            │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐         │
│  │  Dashboard   │  │ Transações   │  │  Categorias  │         │
│  └──────────────┘  └──────────────┘  └──────────────┘         │
│  ┌──────────────┐  ┌──────────────┐                            │
│  │    Metas     │  │  Orçamentos  │                            │
│  └──────────────┘  └──────────────┘                            │
└────────────────────────┬────────────────────────────────────────┘
                         │ HTTP/REST + JWT Token
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│                    BACKEND API (Go)                              │
│                                                                   │
│  ┌───────────────────────────────────────────────────────────┐ │
│  │                    MIDDLEWARE LAYER                        │ │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐               │ │
│  │  │   Auth   │→ │  Logger  │→ │  Error   │               │ │
│  │  │ (JWT)    │  │          │  │ Handler  │               │ │
│  │  └──────────┘  └──────────┘  └──────────┘               │ │
│  └───────────────────────────────────────────────────────────┘ │
│                                                                   │
│  ┌───────────────────────────────────────────────────────────┐ │
│  │                    HANDLER LAYER (Controllers)             │ │
│  │  ┌────────────┐ ┌────────────┐ ┌────────────┐           │ │
│  │  │ Category   │ │Transaction │ │    Goal    │           │ │
│  │  │ Handler    │ │  Handler   │ │  Handler   │           │ │
│  │  └────────────┘ └────────────┘ └────────────┘           │ │
│  │  ┌────────────┐ ┌────────────┐                           │ │
│  │  │   Budget   │ │ Dashboard  │                           │ │
│  │  │  Handler   │ │  Handler   │                           │ │
│  │  └────────────┘ └────────────┘                           │ │
│  └───────────────────────────────────────────────────────────┘ │
│                                                                   │
│  ┌───────────────────────────────────────────────────────────┐ │
│  │                REPOSITORY LAYER (Data Access)              │ │
│  │  ┌────────────┐ ┌────────────┐ ┌────────────┐           │ │
│  │  │ Category   │ │Transaction │ │    Goal    │           │ │
│  │  │Repository  │ │ Repository │ │ Repository │           │ │
│  │  └────────────┘ └────────────┘ └────────────┘           │ │
│  │  ┌────────────┐ ┌────────────┐                           │ │
│  │  │   Budget   │ │ Dashboard  │                           │ │
│  │  │Repository  │ │ Repository │                           │ │
│  │  └────────────┘ └────────────┘                           │ │
│  └───────────────────────────────────────────────────────────┘ │
└────────────────────────┬────────────────────────────────────────┘
                         │ SQL Queries
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│                POSTGRESQL DATABASE (Supabase)                    │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐         │
│  │ transactions │  │  categories  │  │financial_goals│        │
│  └──────────────┘  └──────────────┘  └──────────────┘         │
│  ┌──────────────┐  ┌──────────────┐                            │
│  │   budgets    │  │  auth.users  │                            │
│  └──────────────┘  └──────────────┘                            │
└─────────────────────────────────────────────────────────────────┘
```

## 🔄 Fluxo de Requisição

### Exemplo: Criar uma transação

```
1. Frontend (Next.js)
   ↓
   POST /api/v1/transactions
   Headers: { Authorization: Bearer <token> }
   Body: { type: "expense", amount: 150.50, ... }

2. Backend - Middleware Layer
   ↓
   auth.go: Valida JWT token do Supabase ✓
   ↓
   logger.go: Log da requisição

3. Backend - Handler Layer
   ↓
   transaction_handler.go:
   - Valida dados de entrada
   - Extrai user_id do token
   - Chama repository

4. Backend - Repository Layer
   ↓
   transaction_repository.go:
   - Monta query SQL
   - Executa no banco
   - Retorna resultado

5. PostgreSQL (Supabase)
   ↓
   INSERT INTO transactions ...
   RETURNING *

6. Response (volta pelo mesmo caminho)
   ↓
   Backend → Frontend
   Status: 201 Created
   Body: { success: true, data: {...} }
```

## 📦 Estrutura de Pastas Detalhada

```
fintrack-api/
│
├── cmd/                          # Entry points da aplicação
│   └── api/
│       └── main.go              # ⭐ Inicia o servidor, configura rotas
│
├── internal/                     # Código interno (não exportável)
│   │
│   ├── config/                  # Configurações
│   │   └── config.go           # Carrega .env, conecta DB
│   │
│   ├── handler/                # Controllers HTTP
│   │   ├── category_handler.go    # CRUD de categorias
│   │   ├── transaction_handler.go # CRUD de transações
│   │   ├── goal_handler.go        # CRUD de metas
│   │   ├── budget_handler.go      # CRUD de orçamentos
│   │   ├── dashboard_handler.go   # Dados do dashboard
│   │   └── health_handler.go      # Health check
│   │
│   ├── middleware/             # Middlewares
│   │   ├── auth.go            # 🔐 Valida JWT do Supabase
│   │   ├── logger.go          # 📝 Logs de requisições
│   │   └── error.go           # ❌ Tratamento de erros
│   │
│   ├── models/                # Estruturas de dados
│   │   ├── category.go       # Modelo de categoria
│   │   ├── transaction.go    # Modelo de transação
│   │   ├── goal.go           # Modelo de meta
│   │   ├── budget.go         # Modelo de orçamento
│   │   ├── dashboard.go      # Modelo de dados do dashboard
│   │   ├── user.go           # Modelo de usuário
│   │   └── response.go       # Modelos de resposta HTTP
│   │
│   └── repository/            # Camada de dados
│       ├── category_repository.go    # SQL queries - categorias
│       ├── transaction_repository.go # SQL queries - transações
│       ├── goal_repository.go        # SQL queries - metas
│       ├── budget_repository.go      # SQL queries - orçamentos
│       └── dashboard_repository.go   # SQL queries - dashboard
│
├── docs/                       # Documentação
│   ├── API.md                 # 📚 Documentação completa da API
│   ├── QUICKSTART.md          # 🚀 Guia de início rápido
│   ├── POSTMAN.md             # 📮 Collection do Postman
│   └── FRONTEND-INTEGRATION.md # 🔗 Integração com frontend
│
├── .env.example               # Exemplo de variáveis de ambiente
├── .gitignore                 # Arquivos ignorados pelo git
├── .air.toml                  # Config do Air (hot reload)
├── Dockerfile                 # Build da imagem Docker
├── docker-compose.yml         # Orquestração Docker
├── Makefile                   # Comandos úteis
├── go.mod                     # Dependências do Go
├── go.sum                     # Checksums das dependências
├── README.md                  # 📖 Documentação principal
└── SETUP.md                   # 🛠️ Guia de setup
```

## 🔌 Endpoints e suas Responsabilidades

```
┌─────────────────────────────────────────────────────────────────┐
│                        API ENDPOINTS                             │
├─────────────────────────────────────────────────────────────────┤
│                                                                   │
│  /health                                                          │
│    GET → Health check (público)                                  │
│                                                                   │
│  /api/v1/dashboard/*                          [Protected]        │
│    GET /stats → Estatísticas gerais                             │
│    GET /expenses-by-category → Gastos por categoria             │
│    GET /monthly-data → Dados mensais                             │
│    GET /daily-data → Dados diários                               │
│    GET /recent-transactions → Últimas transações                 │
│                                                                   │
│  /api/v1/categories/*                         [Protected]        │
│    POST   / → Criar categoria                                    │
│    GET    / → Listar categorias                                  │
│    GET    /:id → Buscar categoria                                │
│    PUT    /:id → Atualizar categoria                             │
│    DELETE /:id → Deletar categoria                               │
│                                                                   │
│  /api/v1/transactions/*                       [Protected]        │
│    POST   / → Criar transação                                    │
│    GET    / → Listar transações (com filtros e paginação)        │
│    GET    /:id → Buscar transação                                │
│    PUT    /:id → Atualizar transação                             │
│    DELETE /:id → Deletar transação                               │
│                                                                   │
│  /api/v1/goals/*                              [Protected]        │
│    POST   / → Criar meta                                         │
│    GET    / → Listar metas                                       │
│    GET    /:id → Buscar meta                                     │
│    PUT    /:id → Atualizar meta                                  │
│    DELETE /:id → Deletar meta                                    │
│    POST   /:id/contribute → Contribuir para meta                 │
│                                                                   │
│  /api/v1/budgets/*                            [Protected]        │
│    POST   / → Criar orçamento                                    │
│    GET    / → Listar orçamentos                                  │
│    GET    /with-spent → Orçamentos com valores gastos           │
│    GET    /:id → Buscar orçamento                                │
│    PUT    /:id → Atualizar orçamento                             │
│    DELETE /:id → Deletar orçamento                               │
│                                                                   │
└─────────────────────────────────────────────────────────────────┘
```

## 🗄️ Modelo de Dados

```
┌──────────────────┐       ┌──────────────────┐
│   auth.users     │       │   categories     │
├──────────────────┤       ├──────────────────┤
│ id (UUID) PK     │◄──┐   │ id (UUID) PK     │
│ email            │   │   │ user_id FK       │───┐
│ created_at       │   │   │ name             │   │
└──────────────────┘   │   │ type             │   │
                       │   │ color            │   │
                       │   │ icon             │   │
                       │   │ created_at       │   │
                       │   └──────────────────┘   │
                       │                          │
                       │   ┌──────────────────┐   │
                       │   │  transactions    │   │
                       │   ├──────────────────┤   │
                       └───┤ user_id FK       │   │
                           │ category_id FK   │◄──┘
                           │ type             │
                           │ amount           │
                           │ description      │
                           │ date             │
                           │ created_at       │
                           │ updated_at       │
                           └──────────────────┘

                       ┌───────────────────┐
                       │ financial_goals   │
                       ├───────────────────┤
                   ┌───┤ user_id FK        │
                   │   │ title             │
                   │   │ target_amount     │
                   │   │ current_amount    │
                   │   │ deadline          │
                   │   │ status            │
                   │   │ created_at        │
                   │   │ updated_at        │
                   │   └───────────────────┘
                   │
                   │   ┌──────────────────┐
                   │   │    budgets       │
                   │   ├──────────────────┤
                   └───┤ user_id FK       │
                       │ category_id FK   │
                       │ amount           │
                       │ month            │
                       │ created_at       │
                       └──────────────────┘
```

## 🔐 Fluxo de Autenticação

```
1. Usuário faz login no Frontend
   ↓
2. Frontend chama Supabase Auth
   ↓
3. Supabase retorna JWT token
   ↓
4. Frontend armazena token
   ↓
5. Frontend faz requisição para Backend API
   Authorization: Bearer <token>
   ↓
6. Backend - Middleware de Autenticação
   - Extrai token do header
   - Valida assinatura usando JWT_SECRET
   - Decodifica payload
   - Extrai user_id
   - Adiciona ao contexto da requisição
   ↓
7. Handler acessa user_id do contexto
   ↓
8. Repository filtra dados por user_id
   ↓
9. Response retorna apenas dados do usuário autenticado
```

## 📊 Padrões de Design Utilizados

### 1. Repository Pattern

```
Handler → Repository → Database
(lógica HTTP) (lógica de dados) (PostgreSQL)
```

### 2. Middleware Pattern

```
Request → Auth → Logger → Error Handler → Handler
```

### 3. Dependency Injection

```
main.go:
  db = ConnectDB()
  repo = NewRepository(db)
  handler = NewHandler(repo)
```

## 🎯 Resumo das Responsabilidades

| Camada         | Responsabilidade                                                        |
| -------------- | ----------------------------------------------------------------------- |
| **Handler**    | Recebe HTTP request, valida entrada, chama repository, retorna response |
| **Repository** | Executa queries no banco, mapeia resultados para models                 |
| **Middleware** | Autenticação, logging, tratamento de erros                              |
| **Model**      | Define estruturas de dados e validações                                 |
| **Config**     | Carrega configurações e conecta ao banco                                |

---

Esta arquitetura garante:

- ✅ Separação de responsabilidades
- ✅ Código testável
- ✅ Fácil manutenção
- ✅ Escalabilidade
- ✅ Reutilização de código
