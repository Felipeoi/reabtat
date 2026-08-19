#!/usr/bin/env sh
# Cria/atualiza contas demo no banco de homologacao.
set -e
ROOT="$(CDPATH= cd -- "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

docker compose -f docker-compose.staging.yml exec -T db psql -U postgres -d resocialization_v2_staging < backend/migrations/seed_default_users.sql
echo "OK - usuarios demo atualizados no staging"
