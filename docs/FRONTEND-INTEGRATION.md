# Guia de Integração Frontend → Backend API

Este guia mostra como integrar o frontend Next.js com a API Go.

## 🔧 Configuração Inicial

### 1. Adicione a variável de ambiente no frontend

Edite o arquivo `.env.local` do Next.js:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

### 2. Crie o cliente da API

Crie o arquivo `src/lib/api-client.ts`:

```typescript
import { createClient } from "@supabase/supabase-js";

const supabase = createClient(
  process.env.NEXT_PUBLIC_SUPABASE_URL!,
  process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY!
);

const API_BASE_URL =
  process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api/v1";

async function getAuthToken(): Promise<string | null> {
  const {
    data: { session },
  } = await supabase.auth.getSession();
  return session?.access_token || null;
}

async function apiRequest<T>(
  endpoint: string,
  options: RequestInit = {}
): Promise<T> {
  const token = await getAuthToken();

  if (!token) {
    throw new Error("Not authenticated");
  }

  const response = await fetch(`${API_BASE_URL}${endpoint}`, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
      ...options.headers,
    },
  });

  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || "Request failed");
  }

  return response.json();
}

export const api = {
  // Categories
  categories: {
    list: (type?: "income" | "expense") =>
      apiRequest<{ success: boolean; data: Category[] }>(
        `/categories${type ? `?type=${type}` : ""}`
      ),

    create: (data: CreateCategoryRequest) =>
      apiRequest<{ success: boolean; data: Category }>("/categories", {
        method: "POST",
        body: JSON.stringify(data),
      }),

    update: (id: string, data: UpdateCategoryRequest) =>
      apiRequest<{ success: boolean }>(`/categories/${id}`, {
        method: "PUT",
        body: JSON.stringify(data),
      }),

    delete: (id: string) =>
      apiRequest<{ success: boolean }>(`/categories/${id}`, {
        method: "DELETE",
      }),
  },

  // Transactions
  transactions: {
    list: (filters?: TransactionFilters) => {
      const params = new URLSearchParams();
      if (filters) {
        Object.entries(filters).forEach(([key, value]) => {
          if (value !== undefined && value !== null) {
            params.append(key, String(value));
          }
        });
      }
      return apiRequest<PaginatedResponse<Transaction>>(
        `/transactions?${params.toString()}`
      );
    },

    get: (id: string) =>
      apiRequest<{ success: boolean; data: Transaction }>(
        `/transactions/${id}`
      ),

    create: (data: CreateTransactionRequest) =>
      apiRequest<{ success: boolean; data: Transaction }>("/transactions", {
        method: "POST",
        body: JSON.stringify(data),
      }),

    update: (id: string, data: UpdateTransactionRequest) =>
      apiRequest<{ success: boolean }>(`/transactions/${id}`, {
        method: "PUT",
        body: JSON.stringify(data),
      }),

    delete: (id: string) =>
      apiRequest<{ success: boolean }>(`/transactions/${id}`, {
        method: "DELETE",
      }),
  },

  // Goals
  goals: {
    list: (status?: string) =>
      apiRequest<{ success: boolean; data: FinancialGoal[] }>(
        `/goals${status ? `?status=${status}` : ""}`
      ),

    create: (data: CreateGoalRequest) =>
      apiRequest<{ success: boolean; data: FinancialGoal }>("/goals", {
        method: "POST",
        body: JSON.stringify(data),
      }),

    contribute: (id: string, amount: number) =>
      apiRequest<{ success: boolean }>(`/goals/${id}/contribute`, {
        method: "POST",
        body: JSON.stringify({ amount }),
      }),

    update: (id: string, data: UpdateGoalRequest) =>
      apiRequest<{ success: boolean }>(`/goals/${id}`, {
        method: "PUT",
        body: JSON.stringify(data),
      }),

    delete: (id: string) =>
      apiRequest<{ success: boolean }>(`/goals/${id}`, {
        method: "DELETE",
      }),
  },

  // Budgets
  budgets: {
    list: (month?: string) =>
      apiRequest<{ success: boolean; data: Budget[] }>(
        `/budgets${month ? `?month=${month}` : ""}`
      ),

    listWithSpent: (month?: string) =>
      apiRequest<{ success: boolean; data: BudgetWithSpent[] }>(
        `/budgets/with-spent${month ? `?month=${month}` : ""}`
      ),

    create: (data: CreateBudgetRequest) =>
      apiRequest<{ success: boolean; data: Budget }>("/budgets", {
        method: "POST",
        body: JSON.stringify(data),
      }),

    update: (id: string, data: UpdateBudgetRequest) =>
      apiRequest<{ success: boolean }>(`/budgets/${id}`, {
        method: "PUT",
        body: JSON.stringify(data),
      }),

    delete: (id: string) =>
      apiRequest<{ success: boolean }>(`/budgets/${id}`, {
        method: "DELETE",
      }),
  },

  // Dashboard
  dashboard: {
    stats: (startDate?: string, endDate?: string) => {
      const params = new URLSearchParams();
      if (startDate) params.append("start_date", startDate);
      if (endDate) params.append("end_date", endDate);
      return apiRequest<{ success: boolean; data: DashboardStats }>(
        `/dashboard/stats?${params.toString()}`
      );
    },

    expensesByCategory: (startDate?: string, endDate?: string) => {
      const params = new URLSearchParams();
      if (startDate) params.append("start_date", startDate);
      if (endDate) params.append("end_date", endDate);
      return apiRequest<{ success: boolean; data: CategoryExpense[] }>(
        `/dashboard/expenses-by-category?${params.toString()}`
      );
    },

    monthlyData: (months: number = 6) =>
      apiRequest<{ success: boolean; data: MonthlyData[] }>(
        `/dashboard/monthly-data?months=${months}`
      ),

    dailyData: (startDate?: string, endDate?: string) => {
      const params = new URLSearchParams();
      if (startDate) params.append("start_date", startDate);
      if (endDate) params.append("end_date", endDate);
      return apiRequest<{ success: boolean; data: DailyData[] }>(
        `/dashboard/daily-data?${params.toString()}`
      );
    },

    recentTransactions: (limit: number = 10) =>
      apiRequest<{ success: boolean; data: Transaction[] }>(
        `/dashboard/recent-transactions?limit=${limit}`
      ),
  },
};

// Types
export interface Category {
  id: string;
  user_id: string;
  name: string;
  type: "income" | "expense";
  color: string;
  icon: string;
  created_at: string;
}

export interface CreateCategoryRequest {
  name: string;
  type: "income" | "expense";
  color: string;
  icon: string;
}

export interface UpdateCategoryRequest {
  name?: string;
  type?: "income" | "expense";
  color?: string;
  icon?: string;
}

export interface Transaction {
  id: string;
  user_id: string;
  category_id: string | null;
  type: "income" | "expense";
  amount: number;
  description: string | null;
  date: string;
  created_at: string;
  updated_at: string;
  category?: Category;
}

export interface CreateTransactionRequest {
  category_id?: string;
  type: "income" | "expense";
  amount: number;
  description?: string;
  date: string;
}

export interface UpdateTransactionRequest {
  category_id?: string;
  type?: "income" | "expense";
  amount?: number;
  description?: string;
  date?: string;
}

export interface TransactionFilters {
  type?: "income" | "expense";
  category_id?: string;
  start_date?: string;
  end_date?: string;
  min_amount?: number;
  max_amount?: number;
  page?: number;
  limit?: number;
}

export interface PaginatedResponse<T> {
  success: boolean;
  data: T[];
  page: number;
  limit: number;
  total_count: number;
  total_pages: number;
}

export interface FinancialGoal {
  id: string;
  user_id: string;
  title: string;
  target_amount: number;
  current_amount: number;
  deadline: string | null;
  status: "active" | "completed" | "cancelled";
  created_at: string;
  updated_at: string;
}

export interface CreateGoalRequest {
  title: string;
  target_amount: number;
  current_amount?: number;
  deadline?: string;
}

export interface UpdateGoalRequest {
  title?: string;
  target_amount?: number;
  current_amount?: number;
  deadline?: string;
  status?: "active" | "completed" | "cancelled";
}

export interface Budget {
  id: string;
  user_id: string;
  category_id: string;
  amount: number;
  month: string;
  created_at: string;
  category?: Category;
}

export interface BudgetWithSpent extends Budget {
  spent: number;
  remaining: number;
  percentage: number;
}

export interface CreateBudgetRequest {
  category_id: string;
  amount: number;
  month: string;
}

export interface UpdateBudgetRequest {
  amount?: number;
  month?: string;
}

export interface DashboardStats {
  totalIncome: number;
  totalExpenses: number;
  balance: number;
  savingsRate: number;
}

export interface CategoryExpense {
  category: string;
  amount: number;
  color: string;
  percentage: number;
}

export interface MonthlyData {
  month: string;
  income: number;
  expenses: number;
}

export interface DailyData {
  date: string;
  income: number;
  expenses: number;
}
```

## 📝 Exemplos de Uso

### Em um componente React

```typescript
"use client";

import { useEffect, useState } from "react";
import { api, Transaction } from "@/lib/api-client";

export function TransactionsList() {
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    async function loadTransactions() {
      try {
        const response = await api.transactions.list({
          page: 1,
          limit: 20,
        });
        setTransactions(response.data);
      } catch (err) {
        setError(
          err instanceof Error ? err.message : "Failed to load transactions"
        );
      } finally {
        setLoading(false);
      }
    }

    loadTransactions();
  }, []);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div>
      {transactions.map((transaction) => (
        <div key={transaction.id}>
          <p>{transaction.description}</p>
          <p>${transaction.amount}</p>
        </div>
      ))}
    </div>
  );
}
```

### Criar uma transação

```typescript
async function handleCreateTransaction(data: CreateTransactionRequest) {
  try {
    const response = await api.transactions.create({
      type: "expense",
      amount: 150.5,
      description: "Almoço",
      date: new Date().toISOString(),
    });

    console.log("Transaction created:", response.data);
    // Recarregar lista ou atualizar estado
  } catch (error) {
    console.error("Error creating transaction:", error);
  }
}
```

### Carregar dashboard

```typescript
export function Dashboard() {
  const [stats, setStats] = useState<DashboardStats | null>(null);

  useEffect(() => {
    async function loadDashboard() {
      const [statsRes, expensesRes, monthlyRes] = await Promise.all([
        api.dashboard.stats(),
        api.dashboard.expensesByCategory(),
        api.dashboard.monthlyData(6),
      ]);

      setStats(statsRes.data);
      // ... processar outros dados
    }

    loadDashboard();
  }, []);

  return (
    <div>
      <h1>Balance: ${stats?.balance}</h1>
      <p>Income: ${stats?.totalIncome}</p>
      <p>Expenses: ${stats?.totalExpenses}</p>
    </div>
  );
}
```

### Com React Query (recomendado)

```typescript
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { api } from "@/lib/api-client";

// Hook para listar transações
export function useTransactions(filters?: TransactionFilters) {
  return useQuery({
    queryKey: ["transactions", filters],
    queryFn: () => api.transactions.list(filters),
  });
}

// Hook para criar transação
export function useCreateTransaction() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: api.transactions.create,
    onSuccess: () => {
      // Invalidar cache para recarregar lista
      queryClient.invalidateQueries({ queryKey: ["transactions"] });
    },
  });
}

// Uso no componente
export function TransactionsPage() {
  const { data, isLoading, error } = useTransactions({ page: 1, limit: 20 });
  const createMutation = useCreateTransaction();

  const handleCreate = async (formData: CreateTransactionRequest) => {
    try {
      await createMutation.mutateAsync(formData);
      // Mostrar sucesso
    } catch (error) {
      // Mostrar erro
    }
  };

  // ... render
}
```

## 🔄 Migração Gradual

Você pode migrar gradualmente do Supabase direto para a API:

### Antes (Supabase direto)

```typescript
const { data, error } = await supabase
  .from("transactions")
  .select("*, category:categories(*)")
  .eq("user_id", userId);
```

### Depois (API Go)

```typescript
const { data } = await api.transactions.list();
```

## 🎯 Benefícios da API

1. **Validação centralizada** - Todas as regras de negócio em um lugar
2. **Performance** - Go é muito mais rápido que JavaScript
3. **Type safety** - Contratos claros entre frontend e backend
4. **Escalabilidade** - Fácil adicionar cache, rate limiting, etc.
5. **Segurança** - Validação de JWT e permissões no backend
6. **Logs** - Monitoramento centralizado de todas as operações

## 📊 Tratamento de Erros

```typescript
try {
  const response = await api.transactions.create(data);
  // Sucesso
} catch (error) {
  if (error instanceof Error) {
    if (error.message.includes("Not authenticated")) {
      // Redirecionar para login
      router.push("/auth/login");
    } else if (error.message.includes("Invalid request")) {
      // Mostrar erro de validação
      toast.error("Dados inválidos");
    } else {
      // Erro genérico
      toast.error("Algo deu errado");
    }
  }
}
```

## ✅ Checklist de Integração

- [ ] Criar arquivo `api-client.ts`
- [ ] Adicionar `NEXT_PUBLIC_API_URL` no `.env.local`
- [ ] Testar autenticação e obtenção de token
- [ ] Migrar categorias para a API
- [ ] Migrar transações para a API
- [ ] Migrar metas para a API
- [ ] Migrar orçamentos para a API
- [ ] Migrar dashboard para a API
- [ ] Adicionar tratamento de erros
- [ ] Adicionar loading states
- [ ] Testar em produção

---

Com isso, você terá uma integração completa e robusta entre o frontend Next.js e o backend Go! 🚀
