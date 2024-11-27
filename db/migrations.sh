#!/bin/bash

set -Cu

echo "~~ Running migrations..."

/go/bin/migrate -database $DATABASE_URL -path db/migrations up

echo "~~ Migrations have been applied."

exec "$@"
