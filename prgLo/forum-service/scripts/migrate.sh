#!/bin/sh

# Ожидание доступности базы данных
echo "Ожидание доступности базы данных..."
while ! nc -z $POSTGRES_HOST $POSTGRES_PORT; do
  sleep 0.1
done
echo "База данных доступна"

# Запуск миграций
echo "Запуск миграций..."
/app/forum-service migrate

echo "Миграции завершены" 