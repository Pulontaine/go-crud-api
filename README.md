# Go CRUD API

Простой REST API на языке Go с использованием PostgreSQL и маршрутизатора Gorilla Mux. Реализованы базовые CRUD-операции для управления пользователями.

## Стек технологий

- Go
- PostgreSQL
- Gorilla Mux
- Docker + Docker Compose

## Установка и запуск

### 1. Клонируйте репозиторий

```bash
git clone https://github.com/Pulontaine/go-crud-api"
```

### 2. Запустите проект через Docker
```bash
docker-compose up --build
```
API будет доступен по адресу: http://localhost:8000. 
База данных будет доступна на порту 5433 на хост-машине.

## Примеры HTTP-запросов

Запросы были выполнены в сервисе POSTMAN, дополнительно проверены в TablePlus.

### 1. Создание пользователя
```bash
POST /users
Content-Type: application/json

{
  "name": "Ivan",
  "email": "Ivan@example.com"
}
```

### 2. Получение всех пользователей
```bash
GET /users
```

### 3. Получение пользователя по ID
```bash
GET /users/1
```

### 4. Обновление пользователя
```bash
PUT /users/1
Content-Type: application/json

{
  "name": "Updated Name",
  "email": "new@example.com"
}
```

### 5. Удаление пользователя
```bash 
DELETE /users/1
```
