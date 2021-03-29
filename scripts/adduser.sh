#!/bin/sh
# This script currently adds a default admin user with default creds
docker-compose exec -T postgres psql -U hivemind -d hivemind -c "INSERT INTO public.\"Operators\"(\"Username\", \"Password\", \"Permission\") VALUES ('admin', 'admin', 1);"

