# Ressocialização — sistema para escritórios de advocacia

Aplicação web para **advogados** e **administradores** acompanharem **reeducandos** (pessoas em cumprimento de pena ou medida semelhante): cadastro com origem e destino desejado, regime prisional, dados do defensor responsável, e **correlação automática de casos** (“matches”) entre reeducandos de escritórios diferentes quando há compatibilidade de critérios — útil para articulação de estratégias, transferências ou contato entre defesas.

O backend é **Go (Echo)**, **PostgreSQL** (schema `resocialization` e tabelas públicas de UF/cidade), **JWT** para autenticação. O frontend é **React + TypeScript (Vite)** com layout próprio (CSS).

## Perfis

| Perfil        | Papel no sistema                                      |
|---------------|--------------------------------------------------------|
| **advogado** (`user`) | Cadastra e edita seus reeducandos, consulta matches. |
| **admin**     | Gestão de usuários do escritório / plataforma.       |

## Funcionalidades (estado atual da API e da interface)

- Cadastro e login (**JWT**).
- **Usuários** (CRUD) — restrito a `admin`.
- **Reeducandos** — regime (`fechado` / `semiaberto` / `aberto`), cidade de origem e de destino, nome e telefone do **advogado responsável**; cada registro fica associado ao usuário logado.
- **UFs e cidades** (dados públicos, com seed IBGE) para montar origem/destino.
- **Matches** — listagem e detalhe de pares compatíveis entre reeducandos de usuários distintos (pontuação e regras no backend).

## Estrutura do repositório

```
resocializationV2-main/
├── backend/
│   ├── cmd/server/          # API HTTP (único binário ativo)
│   ├── internal/
│   │   ├── collections/     # Módulos por domínio: auth, user, inmates, ufs, cities, match
│   │   └── entity/        # Tipos compartilhados
│   ├── pkg/                 # config, db (pgx), jwt, logger
│   ├── migrations/          # SQL em ordem + seeds de cidades
│   └── bruno/               # Coleção de requisições (API)
├── frontend/
│   └── src/
│       ├── pages/           # Login, cadastro, usuários, reeducandos, matches
│       ├── components/
│       └── api.ts / types.ts
├── docker-compose.yml       # backend + frontend + postgres + redis
└── docker-compose.local.yml # só Postgres + Redis (dev local do binário Go)
```

Cada coleção em `internal/collections/<nome>/` segue o padrão **delivery (HTTP)** → **usecase** → **repository**.

## Banco de dados e migrations

SQL puro com `psql`, executado a partir de `backend/`:

```bash
cd backend

export DB_DSN="postgres://postgres:1212@127.0.0.1:5433/resocialization_v2?sslmode=disable"

psql "$DB_DSN" -f migrations/01_init.sql
psql "$DB_DSN" -f migrations/02_inmates.up.sql
psql "$DB_DSN" -f migrations/03_match.up.sql
psql "$DB_DSN" -f migrations/04_add_role_to_users.up.sql
psql "$DB_DSN" -f migrations/04_add_status_and_phone_to_users.up.sql
psql "$DB_DSN" -f migrations/seeds_public.sql
psql "$DB_DSN" -f migrations/seed_cities.sql
psql "$DB_DSN" -f migrations/seed_default_users.sql
```

Os seeds de **UFs e cidades** são obrigatórios para telas de origem/destino.

Para reverter migrations parciais (quando existirem `.down.sql`):

```bash
psql "$DB_DSN" -f migrations/04_add_role_to_users.down.sql
psql "$DB_DSN" -f migrations/04_add_status_and_phone_to_users.down.sql
```

**Nota:** a migration `01_init.sql` ainda cria a tabela legada `resocialization.declarations` (contexto antigo do projeto). Não há rotas HTTP ativas para declarações; o produto em uso é reeducandos + matches.

### Usuários de demonstração (após `seed_default_users.sql`)

| E-mail              | Senha          | Perfil   |
|---------------------|----------------|----------|
| `admin@local.dev`   | `Advogado@123` | admin    |
| `advogado@local.dev`| `Advogado@123` | advogado |

Use **somente em ambiente de desenvolvimento**; altere ou remova em produção.

**Importante:** se no banco **não** existirem esses e-mails (por exemplo só há contas que você criou no cadastro), o login com `admin@local.dev` / `advogado@local.dev` sempre falhará até rodar o seed. Com Docker Compose na raiz do projeto:

```bash
sh scripts/seed-demo-users.sh
```

Ou, com o Postgres acessível em `127.0.0.1:5433`: `psql "$DB_DSN" -f backend/migrations/seed_default_users.sql` (a partir da raiz, ajuste o caminho se estiver em `backend/`).

### Cadastro novo (“Criar conta”)

O e-mail é normalizado em **minúsculas**. Se o cadastro parecia falhar mas a conta foi criada, o problema costumava ser o **login automático** logo após o signup: a API devolvia erro em JSON no formato que o front não lia (mensagem genérica “Erro desconhecido”). Isso foi corrigido no cliente e nas respostas de `/api/auth/login` e `/api/auth/signup`.

### Login retorna 401 (“invalid credentials” / e-mail ou senha incorretos)

1. **Confirme e-mail e senha** — para as contas demo, a senha documentada é exatamente `Advogado@123` (maiúscula **A**, arroba, números).
2. **Rode de novo o seed** se você já tinha criado usuário com o mesmo e-mail e outra senha: `psql "$DB_DSN" -f migrations/seed_default_users.sql` — o script agora **atualiza** senha e status em caso de conflito.
3. **Recompile/reinicie o backend** após atualizar o código (imagem Docker antiga continua servindo a versão anterior).
4. O login busca o e-mail **sem diferenciar maiúsculas/minúsculas** e ignorando espaços nas bordas do e-mail gravado no banco.

5. Confirme que o backend em execução é o atual: `curl -s http://localhost:8080/ | head -c 200` deve mostrar `"version":"1.1.0"`. Se aparecer `1.0.0`, pare o processo antigo na porta 8080 e suba o servidor de novo (`docker compose up --build` ou `go run ./cmd/server`).
6. **Use o mesmo e-mail/senha que existem no Postgres** — contas demo só aparecem depois do `seed_default_users.sql` (veja script acima).

## Docker

Pré-requisitos: Docker e plugin Compose.

```bash
docker compose up --build
```

Antes da primeira subida, crie `backend/dist/.env` (não versionado) se ainda não existir: copie `backend/.env.exemple` para `backend/dist/.env` e ajuste `DB_DSN` para **`...@db:5432/resocialization_v2?sslmode=disable`** (host `db` é o serviço no compose).

**Após subir o Postgres pela primeira vez** (volume novo ou banco vazio), aplique migrations e seeds de uma vez:

```bash
sh scripts/apply-migrations-docker.sh
```

Sem isso, telas de reeducandos e UFs falham (`relation ... does not exist` / 500 em `/api/ufs`). Para **só** recriar usuários demo depois: `sh scripts/seed-demo-users.sh`.

- O Postgres do compose é exposto na porta **5433** no host (para não colidir com um PostgreSQL local na 5432).
- **Nome do container:** `resocializationv2-postgres`. **Nome do banco de dados:** `resocialization_v2` (o *schema* SQL das tabelas continua `resocialization`, como nas migrations).
- **Volume Docker:** `resocializationv2_pgdata` — distinto de outros projetos que usavam `postgres-data`. Na primeira subida o Postgres cria o banco `resocialization_v2`. Se você usava o volume antigo e precisa dos dados, faça dump/restore ou mantenha o DSN apontando para a instância antiga; para começar limpo: `docker compose down -v` e suba de novo.
- O Redis do compose usa o container **`resocializationv2-redis`**.
- O serviço `backend` lê variáveis de `backend/dist/.env` (arquivo **gitignored**). No Docker, use `DB_DSN` com host `db` e banco **`resocialization_v2`**, além de `JWT_SECRET` e `FRONTEND_URL`. O Redis no compose é para evoluções futuras; a API atual não depende dele.
- O **frontend** no Docker é buildado com `VITE_API_URL=http://localhost:8080` (navegador no host chama a API publicada na 8080). Erros da API vêm como JSON `{ "error": "..." }` (handler global + login).
- Com o stack no ar, prefira **`sh scripts/apply-migrations-docker.sh`** (aplica tudo no serviço `db`). Alternativa manual: migrations via `psql` em `127.0.0.1:5433` como na seção “Banco de dados e migrations”.
- Para **só** atualizar contas demo: `sh scripts/seed-demo-users.sh`

## Desenvolvimento sem Docker (apenas processos locais)

**Backend**

```bash
cd backend
cp .env.exemple .env   # ajuste DB_DSN (ex.: 127.0.0.1:5433) e FRONTEND_URL se necessário
go mod tidy
go run ./cmd/server
```

**Frontend**

```bash
cd frontend
npm ci
npm run dev
```

Recomenda-se **Node 20+** (Vite 7).

## Testes de API (Bruno)

Abra a pasta `backend/bruno` no [Bruno](https://www.usebruno.com/), ambiente **Local**, e execute **Auth → Inmates / Users** (matches via UI ou HTTP direto).

## Licença e origem

Projeto derivado de um repositório interno de ressocialização; ajuste de licença conforme a política do seu escritório ou órgão.
