#!/bin/bash
set -e

echo "⏳ Iniciando execução das migrations..."

TABLE_COUNT=$(mysql -h db -u projectuser -pprojectpass -N -e "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema='projectgo';")

if [ "$TABLE_COUNT" -eq 0 ]; then
  for f in /migrations/*.sql; do
    echo "➡️ Executando migration: $f"
    mysql -h db -u projectuser -pprojectpass projectgo < "$f"
  done
  echo "✅ Todas as migrations foram aplicadas com sucesso!"
else
  echo "⚠️ Banco já inicializado, pulando migrations."
fi
