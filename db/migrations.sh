#!/bin/bash

echo '~~~ Running migrations... ~~~'

/go/bin/migrate -database $DATABASE_URL -path db/migrations up

echo '~~~ Migrations completed! ~~~'
