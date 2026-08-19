#!/usr/bin/env sh
# Aplica migrations + seeds no banco de homologacao do docker-compose.staging.yml.
# Uso: na raiz do repositorio, com `docker compose -f docker-compose.staging.yml up -d` no ar.
set -e
ROOT="$(CDPATH= cd -- "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

DB_USER=postgres
DB_NAME=resocialization_v2_staging
COMPOSE_FILE=docker-compose.staging.yml

run_sql() {
  rel="$1"
  echo "==> $rel"
  docker compose -f "$COMPOSE_FILE" exec -T db psql -U "$DB_USER" -d "$DB_NAME" -v ON_ERROR_STOP=1 < "$ROOT/$rel"
}

run_sql backend/migrations/01_init.sql
run_sql backend/migrations/02_inmates.up.sql
run_sql backend/migrations/03_match.up.sql
run_sql backend/migrations/04_add_role_to_users.up.sql
run_sql backend/migrations/04_add_status_and_phone_to_users.up.sql
run_sql backend/migrations/seeds_public.sql

run_sql backend/migrations/05_prison_units.up.sql
run_sql backend/migrations/seed_prison_units.sql
run_sql backend/migrations/06_inmates_prison_units.up.sql
run_sql backend/migrations/07_inmate_destination_units.up.sql
run_sql backend/migrations/08_add_oab_cpf_to_users.up.sql
run_sql backend/migrations/09_password_reset_tokens.up.sql
run_sql backend/migrations/10_match_notifications.up.sql
run_sql backend/migrations/11_match_credits_payments.up.sql

run_sql backend/migrations/seed_cities.sql
run_sql backend/migrations/seed_default_users.sql

echo "OK - migrations e seeds aplicados em $DB_NAME"
