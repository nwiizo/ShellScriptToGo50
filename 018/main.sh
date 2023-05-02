#!/bin/bash

config_file="$1"

pgsql_host=$(grep 'host:' "$config_file" | awk '{print $2}')
pgsql_user=$(grep 'user:' "$config_file" | awk '{print $2}')
pgsql_password=$(grep 'password:' "$config_file" | awk '{print $2}')
pgsql_db=$(grep 'dbname:' "$config_file" | awk '{print $2}')

export PGPASSWORD="$pgsql_password"

psql -h "$pgsql_host" -U "$pgsql_user" -d "$pgsql_db" -c "SELECT 1" > /dev/null 2>&1

if [ $? -eq 0 ]; then
  echo "Connection to PostgreSQL server at $pgsql_host is successful."
else
  echo "Connection to PostgreSQL server at $pgsql_host failed."
fi
