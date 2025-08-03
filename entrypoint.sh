#!/bin/sh
set -e

echo "Running migration..."
/app/app db migrate

echo "Starting server..."
exec /app/app server start
