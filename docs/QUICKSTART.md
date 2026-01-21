# 🚀 Guia de Início Rápido - FinTrack Backend API

Este guia te ajudará a configurar e executar a API FinTrack em minutos.

## ⚡ Setup Rápido

### 1. Pré-requisitos

Certifique-se de ter instalado:

- ✅ Go 1.23 ou superior → [Instalar Go](https://go.dev/doc/install)
- ✅ Conta no Supabase → [Criar conta](https://supabase.com)
- ✅ Git

### 2. Clone e Configure

```bash
# Clone o repositório
cd fintrack-go

# Copie o arquivo de exemplo
cp .env.example .env

# Instale as dependências
go mod download
```

### 3. Configure o Supabase

1. Acesse seu projeto no [Supabase Dashboard](https://app.supabase.com)
2. Vá em **Settings** → **API**
3. Copie as seguintes informações:
   - **Project URL**
   - **anon/public key**
   - **service_role key**

4. Vá em **Settings** → **Database**
   - Copie a **Connection String** (modo direto)

5. Vá em **Settings** → **API** → **JWT Settings**
   - Copie o **JWT Secret**

### 4. Configure o arquivo .env

Edite o arquivo `.env` com suas credenciais:

```env
# Server
PORT=8080
ENV=development
API_VERSION=v1

# Supabase - SUBSTITUA COM SUAS CREDENCIAIS
SUPABASE_URL=https://seu-projeto.supabase.co
SUPABASE_ANON_KEY=sua-anon-key-aqui
SUPABASE_SERVICE_ROLE_KEY=sua-service-role-key-aqui
SUPABASE_JWT_SECRET=seu-jwt-secret-aqui

# Database - SUBSTITUA COM SUAS CREDENCIAIS
DB_HOST=db.seu-projeto.supabase.co
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua-senha-do-banco
DB_NAME=postgres
DB_SSLMODE=require

# Frontend
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001
```

### 5. Execute as Migrações do Banco

No painel do Supabase:

1. Vá em **SQL Editor**
2. Execute os scripts da pasta `../fintrackdev/src/scripts/` na ordem:
   - `001_create_tables.sql`
   - `002_create_default_categories.sql`
   - `003_create_update_trigger.sql`

### 6. Execute a API

```bash
# Executar normalmente
go run cmd/api/main.go

# Ou com Make
make run
```

Você verá:

```
✓ Database connection established
🚀 Server starting on port 8080 (env: development)
📚 API documentation available at http://localhost:8080/api/v1
```

### 7. Teste a API

Abra outro terminal e teste:

```bash
# Health check
curl http://localhost:8080/health

# Resposta esperada:
{
  "status": "ok",
  "timestamp": "2025-12-13T...",
  "version": "1.0.0"
}
```

## 🎯 Próximos Passos

### Teste com autenticação

1. **Faça login no frontend** (ou use o Supabase para obter um token)
2. **Copie o token JWT** que o Supabase retorna
3. **Teste um endpoint protegido**:

```bash
# Substitua <SEU-TOKEN> pelo token real
curl -X GET http://localhost:8080/api/v1/categories \
  -H "Authorization: Bearer <SEU-TOKEN>"
```

### Criar sua primeira categoria

```bash
curl -X POST http://localhost:8080/api/v1/categories \
  -H "Authorization: Bearer <SEU-TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Salário",
    "type": "income",
    "color": "#10b981",
    "icon": "dollar-sign"
  }'
```

### Criar sua primeira transação

```bash
curl -X POST http://localhost:8080/api/v1/transactions \
  -H "Authorization: Bearer <SEU-TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "income",
    "amount": 5000.00,
    "description": "Salário de Dezembro",
    "date": "2025-12-13T00:00:00Z"
  }'
```

## 🐳 Executar com Docker (Alternativa)

Se preferir usar Docker:

```bash
# Build da imagem
docker build -t fintrack-api .

# Execute
docker run -p 8080:8080 --env-file .env fintrack-api
```

Ou com Docker Compose:

```bash
docker-compose up
```

## 🔥 Hot Reload (Desenvolvimento)

Para desenvolvimento com recarga automática:

```bash
# Instale o Air
go install github.com/cosmtrek/air@latest

# Execute
air

# Ou com Make
make dev
```

Agora a API será reiniciada automaticamente quando você editar o código!

## 📚 Documentação

- **API Completa**: [docs/API.md](API.md)
- **README Principal**: [../README.md](../README.md)

## 🛠️ Comandos Úteis

```bash
# Ver todos os comandos disponíveis
make help

# Rodar testes
make test

# Build para produção
make build

# Limpar arquivos temporários
make clean

# Ver logs do Docker
make docker-logs
```

## ❓ Problemas Comuns

### Erro de conexão com o banco

```
Failed to connect to database: dial tcp: lookup db.xxx.supabase.co: no such host
```

**Solução**: Verifique se o `DB_HOST` no `.env` está correto.

### Erro de autenticação

```
Invalid or expired token
```

**Solução**:

1. Verifique se o `SUPABASE_JWT_SECRET` está correto
2. Certifique-se de que o token não expirou
3. Faça login novamente para obter um novo token

### Porta já em uso

```
Failed to start server: listen tcp :8080: bind: address already in use
```

**Solução**:

```bash
# Mude a porta no .env
PORT=8081

# Ou mate o processo que está usando a porta 8080
# Windows PowerShell:
Get-Process -Id (Get-NetTCPConnection -LocalPort 8080).OwningProcess | Stop-Process

# Linux/Mac:
lsof -ti:8080 | xargs kill
```

## ✅ Checklist de Configuração

- [ ] Go instalado (versão 1.23+)
- [ ] Repositório clonado
- [ ] Arquivo `.env` criado e configurado
- [ ] Credenciais do Supabase adicionadas
- [ ] Migrações executadas no banco
- [ ] Dependências instaladas (`go mod download`)
- [ ] API rodando (`make run`)
- [ ] Health check funcionando
- [ ] Token JWT obtido
- [ ] Endpoint protegido testado com sucesso

## 🎉 Pronto!

Sua API está configurada e funcionando!

Agora você pode:

- 👉 Integrar com o frontend Next.js
- 👉 Criar categorias, transações, metas e orçamentos
- 👉 Desenvolver novas funcionalidades
- 👉 Testar com o Postman ou Insomnia

## 📞 Suporte

Se tiver problemas:

1. Confira a [documentação completa](API.md)
2. Verifique os logs da aplicação
3. Abra uma issue no GitHub

---

**Feliz codificação! 🚀**
