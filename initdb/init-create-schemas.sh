#!/bin/bash
set -e

# Nama database utama dari environment POSTGRES_DB
DB_NAME="${POSTGRES_DB:-postgres}"  # fallback ke 'postgres' kalau tidak diset

# Pastikan database ada
echo "Ensuring database exists: $DB_NAME"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE DATABASE "$DB_NAME" WITH OWNER "$POSTGRES_USER";
EOSQL

# Setelah database dibuat, buat schemas di dalamnya
echo "Creating schemas: public, n8n in database: $DB_NAME"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname="$DB_NAME" <<-EOSQL
    CREATE SCHEMA IF NOT EXISTS public;
    CREATE SCHEMA IF NOT EXISTS n8n;
EOSQL
