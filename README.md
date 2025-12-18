```markdown
# 🚀 User Management API (Go + Fiber)

Welcome to the **User Management Service**! This is a high-performance REST API built with Go, designed to handle user data efficiently with proper validation, logging, and persistence.

I built this project to demonstrate a production-ready backend architecture using **Golang** and **Docker**. It moves beyond simple CRUD by implementing **pagination**, **structured logging**, and **database migrations**.

---

## 🛠️ Tech Stack & Features
This isn't just a "Hello World" app. It uses industry-standard tools:

* **Language:** [Go 1.25](https://go.dev/) (Latest stable)
* **Framework:** [Fiber v2](https://gofiber.io/) (Fastest HTTP engine for Go)
* **Database:** PostgreSQL 15
* **ORM/SQL:** SQLC (Type-safe SQL generation—no runtime crashes!)
* **Infrastructure:** Docker & Docker Compose
* **Logging:** Uber Zap (Structured, high-performance logging)
* **Validation:** Go-Playground Validator

---

## ⚡ Quick Start (The Easy Way)
The easiest way to run this application is with Docker. You don't need to install Go or Postgres on your machine.

### 1. Clone the repository
```bash
git clone [https://github.com/faisal-990/go_assignment_age_api.git](https://github.com/faisal-990/go_assignment_age_api.git)
cd go_assignment_age_api

```

### 2. Configure Environment

I've included an example config file. Simply copy it to create your local secrets file.

```bash
cp .env.example .env

```

*(You can leave the default values in `.env` as-is for local development)*

### 3. Launch!

Run this command to spin up both the **API** and the **Postgres Database**:

```bash
docker compose up --build

```

That's it! The API is now running at `http://localhost:8080`.

---

## 🐢 Manual Setup (The Hard Way)

If you prefer running it "bare metal" on your machine (without Docker), follow these steps:

**Prerequisites:** You need Go 1.25+ and a running Postgres instance.

1. **Install Dependencies:**
```bash
go mod download

```


2. **Setup Database:**
Create a database named `usermgmt` in your local Postgres.
3. **Run Migrations:**
(Requires `golang-migrate` tool)
```bash
migrate -path db/migrations -database "postgres://user:pass@localhost:5432/usermgmt?sslmode=disable" up

```


4. **Run the Server:**
```bash
go run cmd/server/main.go

```



---

## 🔌 API Endpoints

You can test these using `curl` or Postman.

| Method | Endpoint | Description | Payload / Query Example |
| --- | --- | --- | --- |
| **POST** | `/users` | Create a new user | `{"name": "Alice", "dob": "2000-01-01"}` |
| **GET** | `/users` | List users (Pagination) | `?page=1&limit=10` |
| **GET** | `/users/:id` | Get a single user by ID | N/A |
| **PUT** | `/users/:id` | Update user details | `{"name": "Alice Updated", "dob": "1999-01-01"}` |
| **DELETE** | `/users/:id` | Remove a user | N/A |

### Example Request

Create a user via terminal:

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Go Developer", "dob": "1998-05-20"}'

```

---

## 📂 Project Structure

I followed the Standard Go Project Layout to keep things organized:

* `cmd/server`: Entry point of the application.
* `internal/handler`: Handles HTTP requests and validation.
* `internal/service`: Business logic (Age calculation, etc.).
* `internal/repository`: Database access layer (SQLC generated code).
* `db/migrations`: SQL files for creating/updating database tables.
* `db/sqlc/query`: Raw SQL queries used by SQLC.

---

## 🤝 Contributing

Feel free to fork this repo and submit a PR if you spot any bugs or have ideas for improvements!

```

```
