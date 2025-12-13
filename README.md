# FinTrack Backend API - Go

API robusta em Go para sistema de gestão financeira pessoal, com autenticação via Supabase, CRUD completo de transações, categorias, metas e orçamentos.

## 🚀 Tecnologias

- **Go 1.23+** - Linguagem de programação
- **Gin** - Framework web HTTP
- **PostgreSQL** - Banco de dados (via Supabase)
- **Supabase** - Autenticação e banco de dados
- **Docker** - Containerização
- **Swagger** - Documentação da API

## 📋 Pré-requisitos

- Go 1.23 ou superior
- PostgreSQL (ou conta Supabase)
- Docker e Docker Compose (opcional)
- Make (opcional, mas recomendado)

## 🔧 Instalação e Configuração

### 1. Clone o repositório

```bash
cd backend-go
```

### 2. Configure as variáveis de ambiente

Copie o arquivo `.env.example` para `.env`:

```bash
cp .env.example .env
```

Edite o arquivo `.env` com suas credenciais do Supabase:

```env
PORT=8080
ENV=development

# Supabase Configuration
SUPABASE_URL=https://your-project-ref.supabase.co
SUPABASE_ANON_KEY=your-anon-key
SUPABASE_SERVICE_ROLE_KEY=your-service-role-key
SUPABASE_JWT_SECRET=your-jwt-secret

# Database Configuration
DB_HOST=db.your-project-ref.supabase.co
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your-database-password
DB_NAME=postgres
DB_SSLMODE=require

# CORS
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001
```

### 3. Instale as dependências

```bash
go mod download
```

ou usando Make:

```bash
make install
```

### 4. Execute as migrações do banco de dados

Execute os scripts SQL da pasta `../fintrackdev/src/scripts/` no seu banco de dados Supabase.

## 🏃 Executando a aplicação

### Desenvolvimento

```bash
go run cmd/api/main.go
```

ou usando Make:

```bash
make run
```

### Com hot reload (usando Air)

```bash
# Instale o Air
go install github.com/cosmtrek/air@latest

# Execute
air
```

ou:

```bash
make dev
```

### Produção

```bash
# Build
make build

# Execute o binário
./bin/api
```

## 🐳 Docker

### Build da imagem

```bash
docker build -t fintrack-api:latest .
```

ou:

```bash
make docker-build
```

### Executar com Docker Compose

```bash
docker-compose up -d
```

ou:

```bash
make docker-up
```

### Parar containers

```bash
make docker-down
```

## 📚 Documentação da API

A documentação completa da API está disponível em:

- **Documentação detalhada**: [docs/API.md](docs/API.md)
- **Postman Collection**: [docs/FinTrack-API.postman_collection.json](docs/FinTrack-API.postman_collection.json)

### Endpoints Principais

#### Health Check

- `GET /health` - Verifica o status da API

#### Dashboard

- `GET /api/v1/dashboard/stats` - Estatísticas gerais
- `GET /api/v1/dashboard/expenses-by-category` - Gastos por categoria
- `GET /api/v1/dashboard/monthly-data` - Dados mensais
- `GET /api/v1/dashboard/daily-data` - Dados diários
- `GET /api/v1/dashboard/recent-transactions` - Transações recentes

#### Categorias

- `POST /api/v1/categories` - Criar categoria
- `GET /api/v1/categories` - Listar categorias
- `GET /api/v1/categories/:id` - Buscar categoria
- `PUT /api/v1/categories/:id` - Atualizar categoria
- `DELETE /api/v1/categories/:id` - Deletar categoria

#### Transações

- `POST /api/v1/transactions` - Criar transação
- `GET /api/v1/transactions` - Listar transações (com filtros e paginação)
- `GET /api/v1/transactions/:id` - Buscar transação
- `PUT /api/v1/transactions/:id` - Atualizar transação
- `DELETE /api/v1/transactions/:id` - Deletar transação

#### Metas Financeiras

- `POST /api/v1/goals` - Criar meta
- `GET /api/v1/goals` - Listar metas
- `GET /api/v1/goals/:id` - Buscar meta
- `PUT /api/v1/goals/:id` - Atualizar meta
- `DELETE /api/v1/goals/:id` - Deletar meta
- `POST /api/v1/goals/:id/contribute` - Contribuir para meta

#### Orçamentos

- `POST /api/v1/budgets` - Criar orçamento
- `GET /api/v1/budgets` - Listar orçamentos
- `GET /api/v1/budgets/with-spent` - Orçamentos com valores gastos
- `GET /api/v1/budgets/:id` - Buscar orçamento
- `PUT /api/v1/budgets/:id` - Atualizar orçamento
- `DELETE /api/v1/budgets/:id` - Deletar orçamento

## 🔐 Autenticação

A API utiliza JWT tokens do Supabase. Todas as rotas protegidas requerem o header:

```
Authorization: Bearer <seu-token-jwt>
```

O token é obtido através do login no Supabase (frontend).

## 🧪 Testes

```bash
# Executar testes
make test

# Testes com coverage
make test-coverage
```

## 📁 Estrutura do Projeto

```
backend-go/
├── cmd/
│   └── api/
│       └── main.go              # Entry point da aplicação
├── internal/
│   ├── config/
│   │   └── config.go            # Configurações da aplicação
│   ├── handler/
│   │   ├── category_handler.go # Handlers de categorias
│   │   ├── transaction_handler.go
│   │   ├── goal_handler.go
│   │   ├── budget_handler.go
│   │   ├── dashboard_handler.go
│   │   └── health_handler.go
│   ├── middleware/
│   │   ├── auth.go             # Middleware de autenticação
│   │   ├── logger.go
│   │   └── error.go
│   ├── models/
│   │   ├── category.go         # Modelos de dados
│   │   ├── transaction.go
│   │   ├── goal.go
│   │   ├── budget.go
│   │   ├── dashboard.go
│   │   ├── user.go
│   │   └── response.go
│   └── repository/
│       ├── category_repository.go  # Camada de acesso a dados
│       ├── transaction_repository.go
│       ├── goal_repository.go
│       ├── budget_repository.go
│       └── dashboard_repository.go
├── docs/                       # Documentação
├── .env.example               # Exemplo de variáveis de ambiente
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 🛠️ Comandos Make

```bash
make help              # Mostra todos os comandos disponíveis
make install           # Instala dependências
make build             # Compila a aplicação
make run               # Executa a aplicação
make dev               # Executa com hot reload
make test              # Executa testes
make test-coverage     # Testes com coverage
make lint              # Executa linter
make fmt               # Formata código
make clean             # Limpa arquivos de build
make docker-build      # Build da imagem Docker
make docker-up         # Inicia containers
make docker-down       # Para containers
```

## 🌐 Integração com Frontend

O frontend (Next.js) deve fazer requisições para esta API. Configure a URL base no frontend:

```typescript
// No seu arquivo de configuração do frontend
const API_BASE_URL =
  process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api/v1";
```

Exemplo de requisição do frontend:

```typescript
const response = await fetch(`${API_BASE_URL}/transactions`, {
  headers: {
    Authorization: `Bearer ${supabaseToken}`,
    "Content-Type": "application/json",
  },
});
```

## 📊 Monitoramento

A API inclui:

- Health check endpoint
- Logging de todas as requisições
- Tratamento de erros centralizado
- CORS configurável
- Connection pooling no banco de dados

## 🤝 Contribuindo

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está sob a licença MIT.

## 👨‍💻 Autor

Gildácio Lopes

## 🐛 Reportar Issues

Se encontrar algum problema, por favor abra uma issue no GitHub.

---

**Status**: ✅ Em desenvolvimento ativo
