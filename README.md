# BijakBudget API

Backend service for the BijakBudget App, built using **Go**, **Gin Framework**, and **PostgreSQL** (via **GORM**).

## 🚀 Tech Stack

* **Language:** [Go (Golang)](https://golang.org/)
* **Web Framework:** [Gin](https://gin-gonic.com/)
* **ORM:** [GORM](https://gorm.io/)
* **Database:** [PostgreSQL](https://www.postgresql.org/)

## 🗂️ Architecture

This project strictly follows the principles of **Clean Architecture** / **N-Tier Architecture**, adhering to SOLID and DRY principles:

* ``models/``: Contains the core domain entities (e.g. `User`, `Transaction`, `Category`) with their respective JSON and GORM tags.
* ``repositories/``: The Data Access Layer. Handles all database and query interactions.
* ``services/``: The Business Logic Layer. All rules and business flows live here.
* ``handlers/``: The HTTP Transport Layer. Responsible for parsing JSON requests, passing data to the services, and returning proper JSON responses.
* ``routes/``: Wires up the HTTP endpoints (routes) and maps them to appropriate handlers.

## 🛠️ Setup & Run

### Prerequisites
* Go 1.25+ installed
* PostgreSQL database installed and running

### 1. Environment Setup
Rename or create a `.env` file in the root directory based on your local database credentials:

```ini
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=bijakbudgetdb
DB_PORT=5432
PORT=8080
```

### 2. Download Dependencies
```bash
go mod tidy
```

### 3. Run the Server
```bash
go run cmd/api/main.go
```
*Note: The database models will be automatically migrated when the server starts.*

## 📡 API Endpoints

The API is served under the `/api/v1` prefix.

### Health Check
* `GET /ping` - Returns a simple "pong" response

### Categories
* `GET    /api/v1/categories` - Fetch all categories
* `GET    /api/v1/categories/:id` - Fetch a specific category
* `POST   /api/v1/categories` - Create a new category
* `PUT    /api/v1/categories/:id` - Update a category
* `DELETE /api/v1/categories/:id` - Delete a category

### Transactions
* `GET    /api/v1/transactions` - Fetch all transactions
* `GET    /api/v1/transactions/:id` - Fetch a specific transaction
* `POST   /api/v1/transactions` - Create a new transaction
* `PUT    /api/v1/transactions/:id` - Update a transaction
* `DELETE /api/v1/transactions/:id` - Delete a transaction

### Users
* `GET    /api/v1/users/:id` - Fetch a specific user
* `POST   /api/v1/users` - Create a new user
