# Guia de Conexão Frontend-Backend

Este documento explica como o frontend se conecta ao backend e como resolver problemas comuns.

## Configuração Atual

### Backend (Go + Echo)
- **Porta**: 8080
- **URL**: http://localhost:8080
- **Endpoints**: /api/*

### Frontend (React + Vite)
- **Porta**: 5173 (padrão do Vite)
- **URL**: http://localhost:5173
- **Proxy configurado**: Sim (vite.config.ts)

## Como Funciona

### Desenvolvimento (sem proxy)
Quando você faz uma requisição diretamente para `http://localhost:8080/api/auth/signup`:
1. O navegador verifica se o domínio é diferente (localhost:5173 → localhost:8080)
2. Por serem portas diferentes, ativa verificação CORS
3. Envia requisição OPTIONS (preflight)
4. Backend precisa responder com headers CORS corretos
5. Se aprovado, envia a requisição POST real

### Desenvolvimento (com proxy - CONFIGURAÇÃO ATUAL)
Com o proxy do Vite configurado, quando você faz uma requisição para `/api/auth/signup`:
1. O Vite intercepta requisições que começam com `/api`
2. Redireciona para `http://localhost:8080/api/auth/signup`
3. Como a requisição vem do servidor Vite (não do navegador), **NÃO há verificação CORS**
4. A resposta volta pelo Vite para o navegador
5. **Muito mais simples e confiável em desenvolvimento!**

## Mudanças Realizadas

### 1. Backend (main.go)
✅ CORS configurado para aceitar múltiplas origens
✅ Headers adicionais permitidos
✅ Credentials habilitado
✅ MaxAge aumentado para 1 hora

### 2. Backend (handlers.go)
✅ Logs detalhados adicionados
✅ Mensagens de erro mais claras
✅ Validação melhorada

### 3. Frontend (vite.config.ts)
✅ Proxy configurado para /api
✅ changeOrigin habilitado
✅ Porta 5173 definida

### 4. Frontend (api.ts)
✅ URL relativa em desenvolvimento (usa proxy)
✅ URL absoluta apenas em produção
✅ Logs de debug adicionados
✅ Tratamento de erros melhorado
✅ Mensagens de erro mais claras

### 5. Frontend (.env)
✅ Variável VITE_API_URL configurada
✅ Arquivo .env.example criado

## Como Testar

### 1. Reinicie o Backend
```bash
# No terminal do backend
# Pare o servidor (Ctrl+C)
# Inicie novamente
go run cmd/server/main.go
```

### 2. Reinicie o Frontend
```bash
# No terminal do frontend
# Pare o servidor (Ctrl+C)
npm run dev
```

### 3. Teste o Cadastro
1. Abra http://localhost:5173/signup
2. Preencha o formulário
3. Abra o Console do navegador (F12)
4. Clique em "Criar conta"

### 4. Verifique os Logs

**Console do navegador (F12 > Console):**
- Deve mostrar: `API Success POST /api/auth/signup: {user_id: X}`
- Se erro, mostrará: `API Error [status] POST /api/auth/signup: mensagem`

**Terminal do backend:**
```
INFO signup: creating user email=test@example.com
INFO signup: user created successfully id=1
```

## Troubleshooting

### Erro: "Erro de conexão com o servidor"
**Causa**: Backend não está rodando ou não está na porta 8080

**Solução**:
1. Verifique se o backend está rodando
2. Confirme que está na porta 8080
3. Teste direto no Bruno: Auth > Signup

### Erro: "CORS policy"
**Causa**: Proxy do Vite não está funcionando OU você está usando URL absoluta

**Solução**:
1. Reinicie o frontend (npm run dev)
2. Verifique se `vite.config.ts` tem o proxy configurado
3. Verifique se `api.ts` usa URL relativa em dev

### Erro: "invalid request body"
**Causa**: Dados enviados estão em formato incorreto

**Solução**:
1. Abra F12 > Network
2. Clique na requisição /signup
3. Veja a aba "Payload" - deve ser JSON válido
4. Verifique se os campos name, email, password estão sendo enviados

### Erro: "name, email and password are required"
**Causa**: Algum campo está vazio

**Solução**:
1. Verifique se todos os campos do formulário estão preenchidos
2. Veja o console se há algum erro no formulário

### Requisição OPTIONS retorna 204, mas POST não é enviado
**Causa**: CORS preflight falhou

**Solução**:
1. Com o proxy configurado, isso NÃO deve acontecer
2. Se acontecer, verifique se você não está usando fetch direto com URL absoluta
3. Sempre use as funções em `api.ts` (apiSignup, apiLogin, etc)

## Verificação Rápida

Execute este checklist:

- [ ] Backend rodando na porta 8080
- [ ] Frontend rodando na porta 5173
- [ ] Arquivo `frontend/.env` existe e tem VITE_API_URL
- [ ] Console do navegador (F12) aberto
- [ ] Terminal do backend visível para ver logs

## Testando com Bruno

Se o frontend continuar com problemas, teste o backend diretamente:

```bash
# No Bruno
1. Abra backend/bruno
2. Execute Auth > Signup
3. Se funcionar = problema está no frontend
4. Se não funcionar = problema está no backend
```

## Logs Úteis

### Backend mostra:
```
INFO signup: creating user email=test@example.com
INFO signup: user created successfully id=1
```
✅ Tudo OK!

### Backend mostra:
```
WARN signup: missing required fields
```
❌ Campos vazios chegando do frontend

### Console do navegador mostra:
```
API Success POST /api/auth/signup: {user_id: 1}
```
✅ Tudo OK!

### Console mostra:
```
Network Error POST /api/auth/signup: Failed to fetch
```
❌ Backend não está respondendo

## Próximos Passos

Se ainda houver problemas:

1. Me mostre o log completo do backend
2. Me mostre o log completo do console do navegador (F12)
3. Me mostre a aba Network (F12 > Network) da requisição que falhou
4. Me diga qual mensagem de erro aparece na tela

## Notas de Produção

Em produção, você precisará:
1. Remover o proxy do Vite (build não usa proxy)
2. Configurar CORS no backend para aceitar o domínio de produção
3. Configurar VITE_API_URL para a URL do backend em produção
4. Considerar usar HTTPS
