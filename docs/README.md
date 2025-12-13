# Financial System - Sistema de Gestão Financeira

Sistema completo de gestão financeira pessoal com backend em Go e frontend em Next.js.

## 📁 Estrutura do Monorepo

```
financial-system/
├── backend-go/          # API REST em Go
│   ├── cmd/            # Entry points da aplicação
│   ├── internal/       # Código interno da aplicação
│   ├── docs/           # Documentação da API
│   └── README.md       # Documentação do backend
│
└── fintrackdev/        # Frontend Next.js
    ├── src/            # Código fonte do frontend
    ├── public/         # Arquivos estáticos
    └── README.md       # Documentação do frontend
```

## 🚀 Tecnologias

### Backend (Go)

- **Go 1.23+** - Linguagem de programação
- **Gin** - Framework web HTTP
- **PostgreSQL** - Banco de dados (via Supabase)
- **Supabase** - Autenticação e banco de dados
- **Docker** - Containerização

### Frontend (Next.js)

- **Next.js 16** - Framework React
- **TypeScript** - Tipagem estática
- **Tailwind CSS** - Estilização
- **Supabase** - Autenticação e banco de dados
- **Recharts** - Gráficos e visualizações

## ✨ Funcionalidades

- ✅ **Autenticação** - Login e registro via Supabase com confirmação de email
- ✅ **Dashboard** - Visão geral das finanças com gráficos e estatísticas
- ✅ **Transações** - CRUD completo com filtros e paginação
- ✅ **Categorias** - Organização de receitas e despesas
- ✅ **Metas Financeiras** - Definição e acompanhamento de objetivos
- ✅ **Orçamentos** - Planejamento mensal por categoria
- 🔜 **Relatórios** - Análises detalhadas (em desenvolvimento)
- 🔜 **Configurações** - Personalização da conta (em desenvolvimento)

## 🔧 Começando

### Pré-requisitos

- Node.js 18+ (para o frontend)
- Go 1.23+ (para o backend)
- Conta no Supabase
- Docker (opcional)

### Setup Rápido

1. **Clone o repositório**

```bash
git clone https://github.com/Gildaciolopes/dev-financial-system.git
cd financial-system
```

2. **Configure o Backend**

```bash
cd backend-go
cp .env.example .env
# Edite o .env com suas credenciais do Supabase
go mod download
go run cmd/api/main.go
```

3. **Configure o Frontend**

```bash
cd fintrackdev
cp .env.example .env.local
# Edite o .env.local com suas credenciais do Supabase
npm install
npm run dev
```

4. **Execute as migrações do banco de dados**
   - Acesse o painel do Supabase
   - Execute os scripts SQL da pasta `fintrackdev/src/scripts/`

### Documentação Detalhada

- [Backend - README completo](backend-go/README.md)
- [Backend - Guia de início rápido](backend-go/docs/QUICKSTART.md)
- [Backend - Documentação da API](backend-go/docs/API.md)
- [Backend - Postman Collection](backend-go/docs/POSTMAN.md)

## 🏗️ Arquitetura

### Backend (Clean Architecture)

```
backend-go/
├── cmd/api/              # Entry point
├── internal/
│   ├── config/          # Configurações
│   ├── models/          # Estruturas de dados
│   ├── repository/      # Acesso ao banco de dados
│   ├── handler/         # Controllers HTTP
│   └── middleware/      # Middleware (auth, logging)
└── docs/                # Documentação
```

### Frontend (Next.js App Router)

```
fintrackdev/
├── src/
│   ├── app/            # Páginas e rotas
│   ├── components/     # Componentes React
│   ├── lib/            # Utilitários e configurações
│   └── types/          # Tipos TypeScript
└── public/             # Arquivos estáticos
```

## 🔐 Autenticação

O sistema usa Supabase para autenticação:

1. **Frontend**: Usuário faz login via Supabase
2. **Token JWT**: Supabase retorna um token JWT
3. **API**: Token é enviado no header `Authorization: Bearer <token>`
4. **Validação**: Backend valida o token usando o JWT Secret do Supabase

## 📊 Fluxo de Dados

```
Frontend (Next.js)
    ↓
    ↓ HTTP Request + JWT Token
    ↓
Backend API (Go)
    ↓
    ↓ Valida JWT
    ↓
PostgreSQL (Supabase)
```

## 🌐 URLs

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **API Docs**: http://localhost:8080/api/v1
- **Health Check**: http://localhost:8080/health

## 📝 Variáveis de Ambiente

### Backend (.env)

```env
PORT=8080
SUPABASE_URL=https://xxx.supabase.co
SUPABASE_JWT_SECRET=xxx
DB_HOST=db.xxx.supabase.co
DB_PASSWORD=xxx
```

### Frontend (.env.local)

```env
NEXT_PUBLIC_SUPABASE_URL=https://xxx.supabase.co
NEXT_PUBLIC_SUPABASE_ANON_KEY=xxx
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

## 🐳 Docker

Execute o backend com Docker:

```bash
cd backend-go
docker-compose up
```

## 🧪 Testes

```bash
# Backend
cd backend-go
make test

# Frontend
cd fintrackdev
npm test
```

## 📚 Scripts Úteis

### Backend

```bash
make run          # Rodar o servidor
make build        # Compilar
make test         # Rodar testes
make docker-up    # Docker compose up
```

### Frontend

```bash
npm run dev       # Modo desenvolvimento
npm run build     # Build para produção
npm run start     # Rodar build de produção
npm run lint      # Linter
```

## 🤝 Contribuindo

1. Fork o projeto
2. Crie uma branch (`git checkout -b feature/NovaFuncionalidade`)
3. Commit suas mudanças (`git commit -m 'Add: Nova funcionalidade'`)
4. Push para a branch (`git push origin feature/NovaFuncionalidade`)
5. Abra um Pull Request

## 🗺️ Roadmap

- [x] Autenticação com Supabase
- [x] CRUD de Transações
- [x] CRUD de Categorias
- [x] CRUD de Metas Financeiras
- [x] CRUD de Orçamentos
- [x] Dashboard com estatísticas
- [ ] Relatórios detalhados
- [ ] Configurações de usuário
- [ ] App Mobile (React Native)
- [ ] Notificações push
- [ ] Exportação de dados (PDF, CSV)
- [ ] Modo offline

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👨‍💻 Autor

**Gildácio Lopes**

- GitHub: [@Gildaciolopes](https://github.com/Gildaciolopes)

## 🐛 Reportar Bugs

Se encontrar algum problema, por favor [abra uma issue](https://github.com/Gildaciolopes/dev-financial-system/issues).

## ⭐ Mostre seu apoio

Se este projeto te ajudou, dê uma ⭐️!

---

**Status do Projeto**: 🚧 Em desenvolvimento ativo

**Última atualização**: Dezembro 2025
