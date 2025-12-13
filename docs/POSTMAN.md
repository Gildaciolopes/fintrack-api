# Postman Collection para FinTrack API

Importe esta collection no Postman para testar todos os endpoints da API.

## Como Usar

1. Abra o Postman
2. Clique em **Import**
3. Copie e cole este conteúdo JSON
4. Configure a variável `{{baseUrl}}` para `http://localhost:8080/api/v1`
5. Configure a variável `{{token}}` com seu JWT token do Supabase

## Collection JSON

```json
{
  "info": {
    "name": "FinTrack API",
    "description": "API completa para gestão financeira pessoal",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "variable": [
    {
      "key": "baseUrl",
      "value": "http://localhost:8080/api/v1",
      "type": "string"
    },
    {
      "key": "token",
      "value": "seu-token-jwt-aqui",
      "type": "string"
    }
  ],
  "item": [
    {
      "name": "Health Check",
      "request": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "http://localhost:8080/health",
          "protocol": "http",
          "host": ["localhost"],
          "port": "8080",
          "path": ["health"]
        }
      }
    },
    {
      "name": "Dashboard",
      "item": [
        {
          "name": "Get Stats",
          "request": {
            "method": "GET",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/dashboard/stats?start_date=2025-01-01&end_date=2025-12-31",
              "host": ["{{baseUrl}}"],
              "path": ["dashboard", "stats"],
              "query": [
                { "key": "start_date", "value": "2025-01-01" },
                { "key": "end_date", "value": "2025-12-31" }
              ]
            }
          }
        },
        {
          "name": "Get Expenses by Category",
          "request": {
            "method": "GET",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/dashboard/expenses-by-category",
              "host": ["{{baseUrl}}"],
              "path": ["dashboard", "expenses-by-category"]
            }
          }
        },
        {
          "name": "Get Monthly Data",
          "request": {
            "method": "GET",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/dashboard/monthly-data?months=6",
              "host": ["{{baseUrl}}"],
              "path": ["dashboard", "monthly-data"],
              "query": [{ "key": "months", "value": "6" }]
            }
          }
        },
        {
          "name": "Get Recent Transactions",
          "request": {
            "method": "GET",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/dashboard/recent-transactions?limit=10",
              "host": ["{{baseUrl}}"],
              "path": ["dashboard", "recent-transactions"],
              "query": [{ "key": "limit", "value": "10" }]
            }
          }
        }
      ]
    },
    {
      "name": "Categories",
      "item": [
        {
          "name": "Create Category",
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              },
              {
                "key": "Content-Type",
                "value": "application/json",
                "type": "text"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"name\": \"Freelance\",\n  \"type\": \"income\",\n  \"color\": \"#10b981\",\n  \"icon\": \"briefcase\"\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/categories",
              "host": ["{{baseUrl}}"],
              "path": ["categories"]
            }
          }
        },
        {
          "name": "List Categories",
          "request": {
            "method": "GET",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/categories?type=expense",
              "host": ["{{baseUrl}}"],
              "path": ["categories"],
              "query": [{ "key": "type", "value": "expense" }]
            }
          }
        },
        {
          "name": "Get Category",
          "request": {
            "method": "GET",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/categories/:id",
              "host": ["{{baseUrl}}"],
              "path": ["categories", ":id"],
              "variable": [{ "key": "id", "value": "category-uuid" }]
            }
          }
        },
        {
          "name": "Update Category",
          "request": {
            "method": "PUT",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              },
              {
                "key": "Content-Type",
                "value": "application/json",
                "type": "text"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"name\": \"Alimentação e Bebidas\",\n  \"color\": \"#059669\"\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/categories/:id",
              "host": ["{{baseUrl}}"],
              "path": ["categories", ":id"],
              "variable": [{ "key": "id", "value": "category-uuid" }]
            }
          }
        },
        {
          "name": "Delete Category",
          "request": {
            "method": "DELETE",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/categories/:id",
              "host": ["{{baseUrl}}"],
              "path": ["categories", ":id"],
              "variable": [{ "key": "id", "value": "category-uuid" }]
            }
          }
        }
      ]
    },
    {
      "name": "Transactions",
      "item": [
        {
          "name": "Create Transaction",
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              },
              {
                "key": "Content-Type",
                "value": "application/json",
                "type": "text"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"type\": \"expense\",\n  \"amount\": 150.50,\n  \"description\": \"Almoço no restaurante\",\n  \"date\": \"2025-12-13T00:00:00Z\"\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/transactions",
              "host": ["{{baseUrl}}"],
              "path": ["transactions"]
            }
          }
        },
        {
          "name": "List Transactions",
          "request": {
            "method": "GET",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/transactions?page=1&limit=20&type=expense",
              "host": ["{{baseUrl}}"],
              "path": ["transactions"],
              "query": [
                { "key": "page", "value": "1" },
                { "key": "limit", "value": "20" },
                { "key": "type", "value": "expense" }
              ]
            }
          }
        },
        {
          "name": "Get Transaction",
          "request": {
            "method": "GET",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/transactions/:id",
              "host": ["{{baseUrl}}"],
              "path": ["transactions", ":id"],
              "variable": [{ "key": "id", "value": "transaction-uuid" }]
            }
          }
        },
        {
          "name": "Update Transaction",
          "request": {
            "method": "PUT",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              },
              {
                "key": "Content-Type",
                "value": "application/json",
                "type": "text"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"amount\": 180.00,\n  \"description\": \"Almoço e sobremesa\"\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/transactions/:id",
              "host": ["{{baseUrl}}"],
              "path": ["transactions", ":id"],
              "variable": [{ "key": "id", "value": "transaction-uuid" }]
            }
          }
        },
        {
          "name": "Delete Transaction",
          "request": {
            "method": "DELETE",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/transactions/:id",
              "host": ["{{baseUrl}}"],
              "path": ["transactions", ":id"],
              "variable": [{ "key": "id", "value": "transaction-uuid" }]
            }
          }
        }
      ]
    },
    {
      "name": "Goals",
      "item": [
        {
          "name": "Create Goal",
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              },
              {
                "key": "Content-Type",
                "value": "application/json",
                "type": "text"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"title\": \"Viagem para Europa\",\n  \"target_amount\": 15000.00,\n  \"current_amount\": 2000.00,\n  \"deadline\": \"2026-06-01T00:00:00Z\"\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/goals",
              "host": ["{{baseUrl}}"],
              "path": ["goals"]
            }
          }
        },
        {
          "name": "List Goals",
          "request": {
            "method": "GET",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/goals?status=active",
              "host": ["{{baseUrl}}"],
              "path": ["goals"],
              "query": [{ "key": "status", "value": "active" }]
            }
          }
        },
        {
          "name": "Contribute to Goal",
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              },
              {
                "key": "Content-Type",
                "value": "application/json",
                "type": "text"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"amount\": 500.00\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/goals/:id/contribute",
              "host": ["{{baseUrl}}"],
              "path": ["goals", ":id", "contribute"],
              "variable": [{ "key": "id", "value": "goal-uuid" }]
            }
          }
        }
      ]
    },
    {
      "name": "Budgets",
      "item": [
        {
          "name": "Create Budget",
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              },
              {
                "key": "Content-Type",
                "value": "application/json",
                "type": "text"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"category_id\": \"category-uuid\",\n  \"amount\": 1500.00,\n  \"month\": \"2025-12-01T00:00:00Z\"\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/budgets",
              "host": ["{{baseUrl}}"],
              "path": ["budgets"]
            }
          }
        },
        {
          "name": "List Budgets",
          "request": {
            "method": "GET",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/budgets?month=2025-12-01",
              "host": ["{{baseUrl}}"],
              "path": ["budgets"],
              "query": [{ "key": "month", "value": "2025-12-01" }]
            }
          }
        },
        {
          "name": "Get Budgets with Spent",
          "request": {
            "method": "GET",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{token}}",
                "type": "text"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/budgets/with-spent?month=2025-12-01",
              "host": ["{{baseUrl}}"],
              "path": ["budgets", "with-spent"],
              "query": [{ "key": "month", "value": "2025-12-01" }]
            }
          }
        }
      ]
    }
  ]
}
```

## Variáveis da Collection

| Variável  | Descrição             | Valor Padrão                   |
| --------- | --------------------- | ------------------------------ |
| `baseUrl` | URL base da API       | `http://localhost:8080/api/v1` |
| `token`   | JWT token do Supabase | (vazio - configure após login) |

## Como Obter o Token

1. Faça login no frontend da aplicação
2. Abra o DevTools do navegador (F12)
3. Vá em Application → Local Storage
4. Procure por `supabase.auth.token`
5. Copie o valor do campo `access_token`
6. Cole na variável `{{token}}` no Postman
