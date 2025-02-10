#!/usr/bin/env sh

#set environment variable
GOOSE_DBSTRING="host=${POSTGRES_HOST} port=${POSTGRES_PORT} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DBNAME}"
if [ "$SSL_MODE" = true ]; then
  GOOSE_DBSTRING="${GOOSE_DBSTRING} sslmode=require"
else
  GOOSE_DBSTRING="${GOOSE_DBSTRING} sslmode=disable"
fi

export GOOSE_DBSTRING
export GOOSE_DRIVER=postgres

#run
cd migrations || exit 1

# Check the command-line argument
case "$1" in
  up)
    goose up
    ;;
  down)
    goose down
    ;;
  create)
    goose create "$2" sql
    ;;
  reset)
    goose reset
    ;;
  refresh)
    goose reset && goose up
    ;;
  *)
    echo "Invalid command. Please use one of the following: up, down, create, reset, refresh"
    exit 1
    ;;
esac
