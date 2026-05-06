#!/usr/bin/env sh
# Cria/atualiza contas demo (senha Advogado@123). Rode na raiz do projeto com o Postgres do compose no ar.
set -e
ROOT="$(CDPATH= cd -- "$(dirname "$0")/.." && pwd)"
cd "$ROOT"
docker compose exec -T db psql -U postgres -d resocialization_v2 < backend/migrations/seed_default_users.sql
echo "OK — use admin@local.dev ou advogado@local.dev com senha Advogado@123"
