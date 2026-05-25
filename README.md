# JWT Authentication Service

REST API сервіс автентифікації на основі JWT токенів, написаний на Go.

## Технології

- **Go** — мова програмування
- **Gin** — веб-фреймворк
- **PostgreSQL** — база даних
- **GORM** — ORM для роботи з БД
- **JWT** — токени автентифікації
- **bcrypt** — хешування паролів

## Структура проєкту
jwt-auth/
├── main.go           # Точка входу, роутер
├── database/
│   └── db.go         # Підключення до PostgreSQL
├── handlers/
│   └── auth.go       # Обробники запитів
├── middleware/
│   └── auth.go       # Перевірка JWT токена
└── models/
└── user.go       # Модель користувача

## Встановлення та запуск

### 1. Клонуй репозиторій
```bash
git clone https://github.com/ghostIB/jwt-auth.git
cd jwt-auth
```

### 2. Встанови залежності
```bash
go mod download
```

### 3. Створи базу даних PostgreSQL
```sql
CREATE DATABASE jwt_auth;
CREATE USER jwt_user WITH PASSWORD '1234';
GRANT ALL PRIVILEGES ON DATABASE jwt_auth TO jwt_user;
```

### 4. Запусти сервер
```bash
go run main.go
```

Сервер запуститься на `http://localhost:8080`

## API Endpoints

### Реєстрація
POST /register
**Body:**
```json
{
  "username": "testuser",
  "password": "12345"
}
```
**Відповідь:**
```json
{
  "message": "Користувача створено успішно!"
}
```

---

### Логін
POST /login
**Body:**
```json
{
  "username": "testuser",
  "password": "12345"
}
```
**Відповідь:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

---

### Профіль (захищений роут)
GET /profile
**Headers:**
Authorization: Bearer <your_token>
**Відповідь:**
```json
{
  "message": "Це захищений роут!",
  "user_id": 1,
  "username": "testuser"
}
```

## Приклади використання

### Реєстрація
```bash
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{"username": "testuser", "password": "12345"}'
```

### Логін
```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"username": "testuser", "password": "12345"}'
```

### Отримання профілю
```bash
curl -X GET http://localhost:8080/profile \
-H "Authorization: Bearer <your_token>"
```

## Як працює автентифікація

1. Користувач реєструється — пароль хешується через **bcrypt** і зберігається в БД
2. При логіні — пароль порівнюється з хешем, якщо вірний — генерується **JWT токен**
3. Токен діє **24 години**
4. Для доступу до захищених роутів токен передається в заголовку `Authorization`
5. **Middleware** перевіряє токен перед кожним захищеним запитом