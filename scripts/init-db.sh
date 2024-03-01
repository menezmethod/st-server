#!/bin/bash
set -e

export PGUSER="$POSTGRES_USER"

function create_database_if_not_exists {
    local db_name=$1
    echo "Checking if database '$db_name' exists..."
    if ! psql -lqt | cut -d \| -f 1 | grep -qw "$db_name"; then
        echo "Database '$db_name' does not exist. Creating..."
        psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
            CREATE DATABASE "$db_name";
EOSQL
    else
        echo "Database '$db_name' already exists. Skipping creation."
    fi
}

create_database_if_not_exists "auth_db"
create_database_if_not_exists "journal_db"
create_database_if_not_exists "discussions_db"

echo "Database creation scripts completed."
