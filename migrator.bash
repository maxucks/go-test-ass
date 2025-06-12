#!/bin/bash
set -o allexport; 
source .env; 
set +o allexport

if [ -z "$GOOSE_DBSTRING" ]; then
  echo "Error: GOOSE_DBSTRING must be set"
  exit 1
fi


if [ -z "$GOOSE_CLICKHOUSE_DBSTRING" ]; then
  echo "Error: GOOSE_CLICKHOUSE_DBSTRING must be set"
  exit 1
fi


GOOSE_DRIVER="postgres"
echo "GOOSE_DRIVER='$GOOSE_DRIVER'"

goose -dir migrations/postgres postgres "$GOOSE_DBSTRING" up
goose -dir migrations/clickhouse clickhouse "$GOOSE_CLICKHOUSE_DBSTRING" up
