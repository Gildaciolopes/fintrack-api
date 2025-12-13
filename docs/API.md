# FinTrack API - Documentação Completa

Esta documentação fornece exemplos detalhados de como usar a API FinTrack.

## Base URL

```
http://localhost:8080/api/v1
```

## Autenticação

Todas as rotas (exceto `/health`) requerem autenticação via JWT token do Supabase.

```
Authorization: Bearer <seu-token-jwt>
```

---

## 🏥 Health Check

### GET /health

Verifica se a API está funcionando.

**Resposta:**

```json
{
  "status": "ok",
  "timestamp": "2025-12-13T10:30:00Z",
  "version": "1.0.0"
}
```

---

## 📊 Dashboard

### GET /api/v1/dashboard/stats

Retorna estatísticas financeiras do usuário.

**Query Parameters:**

- `start_date` (opcional): Data inicial no formato YYYY-MM-DD
- `end_date` (opcional): Data final no formato YYYY-MM-DD

**Exemplo:**

```bash
curl -X GET "http://localhost:8080/api/v1/dashboard/stats?start_date=2025-01-01&end_date=2025-12-31" \
  -H "Authorization: Bearer <token>"
```

**Resposta:**

```json
{
  "success": true,
  "data": {
    "totalIncome": 15000.5,
    "totalExpenses": 8500.75,
    "balance": 6499.75,
    "savingsRate": 43.33
  }
}
```

### GET /api/v1/dashboard/expenses-by-category

Retorna gastos agrupados por categoria.

**Query Parameters:**

- `start_date` (opcional): Data inicial
- `end_date` (opcional): Data final

**Resposta:**

```json
{
  "success": true,
  "data": [
    {
      "category": "Alimentação",
      "amount": 2500.0,
      "color": "#10b981",
      "percentage": 29.41
    },
    {
      "category": "Transporte",
      "amount": 1800.0,
      "color": "#3b82f6",
      "percentage": 21.18
    }
  ]
}
```

### GET /api/v1/dashboard/monthly-data

Retorna dados mensais de receitas e despesas.

**Query Parameters:**

- `months` (opcional): Número de meses (1-12, padrão: 6)

**Resposta:**

```json
{
  "success": true,
  "data": [
    {
      "month": "2025-06",
      "income": 5000.0,
      "expenses": 3200.0
    },
    {
      "month": "2025-07",
      "income": 5200.0,
      "expenses": 3500.0
    }
  ]
}
```

### GET /api/v1/dashboard/daily-data

Retorna dados diários de transações.

**Query Parameters:**

- `start_date` (opcional): Data inicial
- `end_date` (opcional): Data final

**Resposta:**

```json
{
  "success": true,
  "data": [
    {
      "date": "2025-12-01",
      "income": 500.0,
      "expenses": 150.0
    },
    {
      "date": "2025-12-02",
      "income": 0.0,
      "expenses": 200.0
    }
  ]
}
```

### GET /api/v1/dashboard/recent-transactions

Retorna transações recentes.

**Query Parameters:**

- `limit` (opcional): Número de transações (1-50, padrão: 10)

**Resposta:**

```json
{
  "success": true,
  "data": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "user_id": "user-uuid",
      "category_id": "cat-uuid",
      "type": "expense",
      "amount": 150.0,
      "description": "Almoço no restaurante",
      "date": "2025-12-13T00:00:00Z",
      "created_at": "2025-12-13T10:00:00Z",
      "updated_at": "2025-12-13T10:00:00Z",
      "category": {
        "id": "cat-uuid",
        "name": "Alimentação",
        "type": "expense",
        "color": "#10b981",
        "icon": "utensils"
      }
    }
  ]
}
```

---

## 📁 Categorias

### POST /api/v1/categories

Cria uma nova categoria.

**Body:**

```json
{
  "name": "Freelance",
  "type": "income",
  "color": "#10b981",
  "icon": "briefcase"
}
```

**Resposta:**

```json
{
  "success": true,
  "message": "Category created successfully",
  "data": {
    "id": "cat-uuid",
    "user_id": "user-uuid",
    "name": "Freelance",
    "type": "income",
    "color": "#10b981",
    "icon": "briefcase",
    "created_at": "2025-12-13T10:00:00Z"
  }
}
```

### GET /api/v1/categories

Lista todas as categorias do usuário.

**Query Parameters:**

- `type` (opcional): Filtrar por tipo (`income` ou `expense`)

**Exemplo:**

```bash
curl -X GET "http://localhost:8080/api/v1/categories?type=expense" \
  -H "Authorization: Bearer <token>"
```

**Resposta:**

```json
{
  "success": true,
  "data": [
    {
      "id": "cat-uuid-1",
      "user_id": "user-uuid",
      "name": "Alimentação",
      "type": "expense",
      "color": "#10b981",
      "icon": "utensils",
      "created_at": "2025-12-01T10:00:00Z"
    },
    {
      "id": "cat-uuid-2",
      "user_id": "user-uuid",
      "name": "Transporte",
      "type": "expense",
      "color": "#3b82f6",
      "icon": "car",
      "created_at": "2025-12-02T10:00:00Z"
    }
  ]
}
```

### GET /api/v1/categories/:id

Busca uma categoria específica.

**Resposta:**

```json
{
  "success": true,
  "data": {
    "id": "cat-uuid",
    "user_id": "user-uuid",
    "name": "Alimentação",
    "type": "expense",
    "color": "#10b981",
    "icon": "utensils",
    "created_at": "2025-12-01T10:00:00Z"
  }
}
```

### PUT /api/v1/categories/:id

Atualiza uma categoria.

**Body:**

```json
{
  "name": "Alimentação e Bebidas",
  "color": "#059669"
}
```

**Resposta:**

```json
{
  "success": true,
  "message": "Category updated successfully"
}
```

### DELETE /api/v1/categories/:id

Deleta uma categoria.

**Resposta:**

```json
{
  "success": true,
  "message": "Category deleted successfully"
}
```

---

## 💰 Transações

### POST /api/v1/transactions

Cria uma nova transação.

**Body:**

```json
{
  "category_id": "cat-uuid",
  "type": "expense",
  "amount": 150.5,
  "description": "Almoço no restaurante",
  "date": "2025-12-13T00:00:00Z"
}
```

**Resposta:**

```json
{
  "success": true,
  "message": "Transaction created successfully",
  "data": {
    "id": "trans-uuid",
    "user_id": "user-uuid",
    "category_id": "cat-uuid",
    "type": "expense",
    "amount": 150.5,
    "description": "Almoço no restaurante",
    "date": "2025-12-13T00:00:00Z",
    "created_at": "2025-12-13T10:00:00Z",
    "updated_at": "2025-12-13T10:00:00Z"
  }
}
```

### GET /api/v1/transactions

Lista transações com filtros e paginação.

**Query Parameters:**

- `type` (opcional): Tipo (`income` ou `expense`)
- `category_id` (opcional): UUID da categoria
- `start_date` (opcional): Data inicial (YYYY-MM-DD)
- `end_date` (opcional): Data final (YYYY-MM-DD)
- `min_amount` (opcional): Valor mínimo
- `max_amount` (opcional): Valor máximo
- `page` (opcional): Número da página (padrão: 1)
- `limit` (opcional): Itens por página (padrão: 20, máx: 100)

**Exemplo:**

```bash
curl -X GET "http://localhost:8080/api/v1/transactions?type=expense&start_date=2025-12-01&limit=10" \
  -H "Authorization: Bearer <token>"
```

**Resposta:**

```json
{
  "success": true,
  "data": [
    {
      "id": "trans-uuid",
      "user_id": "user-uuid",
      "category_id": "cat-uuid",
      "type": "expense",
      "amount": 150.5,
      "description": "Almoço",
      "date": "2025-12-13T00:00:00Z",
      "created_at": "2025-12-13T10:00:00Z",
      "updated_at": "2025-12-13T10:00:00Z",
      "category": {
        "id": "cat-uuid",
        "name": "Alimentação",
        "type": "expense",
        "color": "#10b981",
        "icon": "utensils"
      }
    }
  ],
  "page": 1,
  "limit": 10,
  "total_count": 45,
  "total_pages": 5
}
```

### GET /api/v1/transactions/:id

Busca uma transação específica.

**Resposta:**

```json
{
  "success": true,
  "data": {
    "id": "trans-uuid",
    "user_id": "user-uuid",
    "category_id": "cat-uuid",
    "type": "expense",
    "amount": 150.5,
    "description": "Almoço",
    "date": "2025-12-13T00:00:00Z",
    "created_at": "2025-12-13T10:00:00Z",
    "updated_at": "2025-12-13T10:00:00Z",
    "category": {
      "id": "cat-uuid",
      "name": "Alimentação",
      "type": "expense",
      "color": "#10b981",
      "icon": "utensils"
    }
  }
}
```

### PUT /api/v1/transactions/:id

Atualiza uma transação.

**Body:**

```json
{
  "amount": 180.0,
  "description": "Almoço e sobremesa"
}
```

**Resposta:**

```json
{
  "success": true,
  "message": "Transaction updated successfully"
}
```

### DELETE /api/v1/transactions/:id

Deleta uma transação.

**Resposta:**

```json
{
  "success": true,
  "message": "Transaction deleted successfully"
}
```

---

## 🎯 Metas Financeiras

### POST /api/v1/goals

Cria uma nova meta financeira.

**Body:**

```json
{
  "title": "Viagem para Europa",
  "target_amount": 15000.0,
  "current_amount": 2000.0,
  "deadline": "2026-06-01T00:00:00Z"
}
```

**Resposta:**

```json
{
  "success": true,
  "message": "Goal created successfully",
  "data": {
    "id": "goal-uuid",
    "user_id": "user-uuid",
    "title": "Viagem para Europa",
    "target_amount": 15000.0,
    "current_amount": 2000.0,
    "deadline": "2026-06-01T00:00:00Z",
    "status": "active",
    "created_at": "2025-12-13T10:00:00Z",
    "updated_at": "2025-12-13T10:00:00Z"
  }
}
```

### GET /api/v1/goals

Lista todas as metas.

**Query Parameters:**

- `status` (opcional): Filtrar por status (`active`, `completed`, `cancelled`)

**Resposta:**

```json
{
  "success": true,
  "data": [
    {
      "id": "goal-uuid",
      "user_id": "user-uuid",
      "title": "Viagem para Europa",
      "target_amount": 15000.0,
      "current_amount": 2000.0,
      "deadline": "2026-06-01T00:00:00Z",
      "status": "active",
      "created_at": "2025-12-13T10:00:00Z",
      "updated_at": "2025-12-13T10:00:00Z"
    }
  ]
}
```

### GET /api/v1/goals/:id

Busca uma meta específica.

### PUT /api/v1/goals/:id

Atualiza uma meta.

**Body:**

```json
{
  "current_amount": 3500.0,
  "status": "active"
}
```

### DELETE /api/v1/goals/:id

Deleta uma meta.

### POST /api/v1/goals/:id/contribute

Adiciona uma contribuição à meta.

**Body:**

```json
{
  "amount": 500.0
}
```

**Resposta:**

```json
{
  "success": true,
  "message": "Contribution added successfully"
}
```

---

## 💵 Orçamentos

### POST /api/v1/budgets

Cria um novo orçamento.

**Body:**

```json
{
  "category_id": "cat-uuid",
  "amount": 1500.0,
  "month": "2025-12-01T00:00:00Z"
}
```

**Resposta:**

```json
{
  "success": true,
  "message": "Budget created successfully",
  "data": {
    "id": "budget-uuid",
    "user_id": "user-uuid",
    "category_id": "cat-uuid",
    "amount": 1500.0,
    "month": "2025-12-01T00:00:00Z",
    "created_at": "2025-12-13T10:00:00Z"
  }
}
```

### GET /api/v1/budgets

Lista todos os orçamentos.

**Query Parameters:**

- `month` (opcional): Filtrar por mês (YYYY-MM-DD)

### GET /api/v1/budgets/with-spent

Lista orçamentos com valores gastos.

**Query Parameters:**

- `month` (opcional): Mês (YYYY-MM-DD, padrão: mês atual)

**Resposta:**

```json
{
  "success": true,
  "data": [
    {
      "id": "budget-uuid",
      "user_id": "user-uuid",
      "category_id": "cat-uuid",
      "amount": 1500.0,
      "month": "2025-12-01T00:00:00Z",
      "created_at": "2025-12-13T10:00:00Z",
      "category": {
        "id": "cat-uuid",
        "name": "Alimentação",
        "type": "expense",
        "color": "#10b981",
        "icon": "utensils"
      },
      "spent": 850.5,
      "remaining": 649.5,
      "percentage": 56.7
    }
  ]
}
```

### GET /api/v1/budgets/:id

Busca um orçamento específico.

### PUT /api/v1/budgets/:id

Atualiza um orçamento.

**Body:**

```json
{
  "amount": 2000.0
}
```

### DELETE /api/v1/budgets/:id

Deleta um orçamento.

---

## ❌ Tratamento de Erros

Todos os endpoints podem retornar os seguintes erros:

### 400 Bad Request

```json
{
  "success": false,
  "error": "Invalid request data",
  "message": "Detalhes do erro de validação"
}
```

### 401 Unauthorized

```json
{
  "success": false,
  "error": "Invalid or expired token",
  "message": "Please login again"
}
```

### 404 Not Found

```json
{
  "success": false,
  "error": "Resource not found",
  "message": "The requested resource was not found"
}
```

### 500 Internal Server Error

```json
{
  "success": false,
  "error": "Internal server error",
  "message": "An unexpected error occurred"
}
```

---

## 🔧 Exemplos de Integração com Frontend

### React/Next.js

```typescript
// lib/api.ts
const API_BASE_URL =
  process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api/v1";

export async function fetchTransactions(token: string, filters?: any) {
  const queryParams = new URLSearchParams(filters).toString();
  const response = await fetch(`${API_BASE_URL}/transactions?${queryParams}`, {
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch transactions");
  }

  return response.json();
}

export async function createTransaction(token: string, data: any) {
  const response = await fetch(`${API_BASE_URL}/transactions`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    throw new Error("Failed to create transaction");
  }

  return response.json();
}
```

### Uso com Supabase

```typescript
import { createClient } from "@supabase/supabase-js";

const supabase = createClient(
  process.env.NEXT_PUBLIC_SUPABASE_URL!,
  process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY!
);

// Obter token do Supabase
const {
  data: { session },
} = await supabase.auth.getSession();
const token = session?.access_token;

// Usar token na API
const transactions = await fetchTransactions(token);
```

---

## 📝 Notas Importantes

1. **Datas**: Todas as datas devem estar no formato ISO 8601 (YYYY-MM-DDTHH:mm:ssZ)
2. **UUIDs**: Todos os IDs são UUIDs v4
3. **Valores monetários**: Sempre em formato decimal com 2 casas decimais
4. **Paginação**: Use os parâmetros `page` e `limit` para controlar a paginação
5. **Rate Limiting**: Considere implementar rate limiting em produção

---

**Versão da API**: 1.0.0  
**Última atualização**: Dezembro 2025
