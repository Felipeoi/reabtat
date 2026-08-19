# Fluxo de Git e ambientes

Este projeto deve ter dois ambientes principais:

- `staging`: base de teste/homologacao.
- `main`: producao.

## Branches

- `main`: somente codigo aprovado para producao.
- `staging`: codigo que esta em teste.
- `codex/<nome>` ou `feature/<nome>`: alteracoes em desenvolvimento.
- `fix/<nome>`: correcoes pontuais.

Fluxo:

```text
feature/fix -> staging -> main
```

## Homologacao local com Docker

Crie o arquivo de variaveis da homologacao:

```bash
cp backend/.env.staging.example backend/.env.staging
```

Suba o ambiente de teste:

```bash
docker compose -f docker-compose.staging.yml up -d --build
```

Aplique migrations e seeds:

```bash
sh scripts/apply-migrations-staging-docker.sh
```

URLs da homologacao local:

- Frontend: http://localhost:25173
- Backend: http://localhost:28080
- Postgres: localhost:25433
- MailHog: http://localhost:28025

Banco da homologacao:

```text
resocialization_v2_staging
```

## Producao

A producao deve usar outro banco e outro arquivo de ambiente, nunca os dados de homologacao:

```text
resocialization_v2_production
```

Checklist antes de subir para producao:

1. Fazer merge da branch de trabalho em `staging`.
2. Subir o ambiente de homologacao.
3. Rodar migrations em homologacao.
4. Testar login, cadastro, reeducandos, matches, pagamentos e recuperacao de senha.
5. Fazer merge de `staging` em `main`.
6. Subir `main` em producao.
7. Rodar migrations no banco de producao.

## Contas demo

Depois dos seeds:

- `admin@local.dev`
- `advogado@local.dev`
- senha: `Advogado@123`
