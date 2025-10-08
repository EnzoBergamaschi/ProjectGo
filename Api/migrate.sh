#!/bin/bash
set -e

echo "⏳ Iniciando execução das migrations..."

for f in /migrations/*.sql; do
  echo "➡️ Executando migration: $f"
  mysql -h db -u projectuser -pprojectpass projectgo < "$f"
done

echo "✅ Todas as migrations foram aplicadas com sucesso!"
