# 🚀 Instruções de Deploy no Render

## ⚠️ IMPORTANTE: Render não está atualizando automaticamente!

O Render está rodando **código antigo** mesmo após commits. Siga os passos abaixo:

## Passos para Forçar Redeploy:

### 1. Verificar último commit local:
```bash
cd fintrack-api
git log --oneline -5
```

Você deve ver:
```
78e3636 fix: corrigir GetMonthlyData usando make_interval()
11a75f7 fix: corrigir queries SQL dos endpoints do dashboard
```

### 2. Acessar Render Dashboard:
https://dashboard.render.com

### 3. Selecionar o serviço:
- Clique em **fintrack-api**

### 4. Verificar branch configurada:
- Vá em **Settings** → **Build & Deploy**
- Confirme que está em **Branch: develop**
- Se estiver em outra branch, mude para **develop** e salve

### 5. Forçar Redeploy Manual:
- Volte para a aba **Events** ou **Deploys**
- Clique no botão **"Manual Deploy"**
- Selecione **"Deploy latest commit"**
- Aguarde o build completar (2-5 minutos)

### 6. Monitorar Logs:
- Clique na aba **"Logs"**
- Aguarde até ver:
  ```
  [GIN-debug] Listening and serving HTTP on :8080
  Server running on port 8080
  ```

### 7. Verificar se os logs de debug aparecem:
Quando recarregar o dashboard, você deve ver nos logs:
```
[DEBUG GetStats] userID: xxx, startDate: 2025-12-01, endDate: 2025-12-31
[DEBUG GetRecentTransactions] userID: xxx, limit: 5
```

## ✅ Verificação Final:

1. **Recarregue o dashboard**: http://localhost:3000/dashboard
2. **Todos os endpoints devem funcionar**:
   - ✅ Stats (receita/despesa/balanço)
   - ✅ Expenses by category (gráfico)
   - ✅ Monthly data (gráfico mensal)
   - ✅ Daily data (gráfico diário)
   - ✅ Recent transactions (tabela)

3. **Teste criar transação sem categoria**:
   - Vá em Transactions → Add Transaction
   - Preencha apenas: Type, Amount, Date
   - Deixe Category e Description vazios
   - Deve funcionar normalmente

## 🐛 Se ainda houver erros:

1. **Verifique logs do Render** para ver exatamente qual erro está acontecendo
2. **Verifique se o commit está correto**: 
   ```bash
   git show 78e3636
   ```
3. **Se necessário, force push**:
   ```bash
   git push -f origin develop
   ```

## 📝 Notas:

- **Auto-deploy**: Render deveria fazer deploy automaticamente após push, mas não está funcionando
- **Solução**: Sempre fazer **Manual Deploy** após push
- **Verificação**: Sempre conferir os logs no Render após deploy para confirmar que está rodando o código correto
