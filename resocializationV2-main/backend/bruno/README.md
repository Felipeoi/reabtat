# Coleção Bruno — API Ressocialização (advocacia)

Testes manuais dos endpoints usados pelo sistema de **reeducandos** e **matches**.

## Pré-requisitos

1. [Bruno](https://www.usebruno.com/)
2. Backend em `http://localhost:8080`
3. PostgreSQL com migrations e seeds aplicados

## Uso

1. **Open Collection** → pasta `backend/bruno`
2. Ambiente **Local** (`baseUrl` = `http://localhost:8080`)

### Ordem sugerida

1. **Auth > Signup** — cria usuário (grava `userId` se o script da requisição estiver configurado)
2. **Auth > Login** — obtém `token`
3. **Inmates** — CRUD de reeducandos (`inmateId` quando aplicável)
4. **Users** — apenas com token de **admin** (CRUD de usuários)

Para **matches**, use o frontend ou `curl` em `/api/matches` e `/api/matches/:myInmateId/:matchedInmateId` com JWT de usuário `user` ou `admin`.

## Variáveis de ambiente (Bruno)

| Variável  | Descrição                    |
|-----------|------------------------------|
| `baseUrl` | URL da API                   |
| `token`   | JWT após login               |
| `userId`  | Último usuário criado (testes) |
| `inmateId`| Último reeducando (testes)   |

## Estrutura

```
bruno/
├── bruno.json
├── environments/Local.bru
├── Auth/
├── Users/
└── Inmates/
```
