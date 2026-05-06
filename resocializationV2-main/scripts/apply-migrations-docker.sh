#!/usr/bin/env sh
# Aplica todas as migrations + seeds no Postgres do docker-compose (serviço db).
# Uso: na raiz do repositório, com `docker compose up -d` (db no ar).
set -e
ROOT="$(CDPATH= cd -- "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

DB_USER=postgres
DB_NAME=resocialization_v2

run_sql() {
  rel="$1"
  echo "==> $rel"
  docker compose exec -T db psql -U "$DB_USER" -d "$DB_NAME" -v ON_ERROR_STOP=1 < "$ROOT/$rel"
}

run_sql backend/migrations/01_init.sql
run_sql backend/migrations/02_inmates.up.sql
run_sql backend/migrations/03_match.up.sql
run_sql backend/migrations/04_add_role_to_users.up.sql
run_sql backend/migrations/04_add_status_and_phone_to_users.up.sql
run_sql backend/migrations/seeds_public.sql

# seed_cities.sql usa IDs fixos: só roda se a tabela estiver vazia (reexecução do script).
cnt="$(docker compose exec -T db psql -U "$DB_USER" -d "$DB_NAME" -t -A -c "SELECT count(*)::bigint FROM public.cities" | tr -d '[:space:]')"
if [ "${cnt:-0}" -eq 0 ]; then
  run_sql backend/migrations/seed_cities.sql
else
  echo "==> backend/migrations/seed_cities.sql (omitido: já existem $cnt cidades)"
fi

run_sql backend/migrations/seed_default_users.sql

echo "OK — migrations e seeds aplicados em $DB_NAME"
