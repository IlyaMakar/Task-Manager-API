# Task Manager API 🚀

Простой и эффективный RESTful API для управления задачами, написанный на Go.  
Этот проект позволяет создавать, просматривать, обновлять и удалять задачи, а также управлять пользователями и аутентификацией.

## 🔹 Особенности

- **Управление задачами** (CRUD операции)
- **Логирование запросов** (middleware)
- **Конфигурация через переменные окружения** (`.env`)
- **Документация API** (Swagger/Postman)

## 🔹 Технологии

- **Язык**: Go (Golang)
- **Фреймворк**: [Gin](https://github.com/gin-gonic/gin) (HTTP-роутер)
- **База данных**: PostgreSQL (через [pgx](https://github.com/jackc/pgx))
- **Конфигурация**: [godotenv](https://github.com/joho/godotenv)


## 🔹 Установка и запуск

### Предварительные требования
- Установленный **Go** (версия 1.20+)
- **PostgreSQL** (или Docker для запуска контейнера)
- Утилита **make** (опционально)

### 1. Клонировать репозиторий
```bash
git clone https://github.com/IlyaMakar/Task-Manager-API.git
cd Task-Manager-API
