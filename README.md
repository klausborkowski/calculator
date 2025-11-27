# Read.me contents

1. [Overview](#1-overview)
2. [Requirements](#2-requirements)
3. [Running the Service](#3-running-the-service)
4. [Check Test Coverage Information](#4-check-test-coverage-information)
5. [UI](#5-ui)
6. [API](#6-api)
7. [Deployed Service](#7-deployed-service)

---

# Packager Service

## 1. Overview  
Our customers can order any number of these items through our website, but they will always only be given complete packs.

### Order Fulfillment Rules:
- Only whole packs can be sent. Packs cannot be broken open.
- Within the constraints of Rule 1 above, send out the least amount of items to fulfill the order.
- Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfill each order.

---

## 2. Requirements

### Docker-based workflows (`make start`, `make docker-up`)
- **Docker Engine** ≥ 29.0.4
- **Docker Compose** ≥ v2.40.3
- **Make** (usually preinstalled on Linux/macOS)

### Hybrid/local workflows (`make start-local`, manual dev)
- **Go** 1.22+
- **Node.js** 18+ and **npm** 9+
- **PostgreSQL** 16+ (can be provided by Docker when using `make start-local`)
- **Make**

---

## 3. Running the Service

### Method 1: Docker Setup (`make start`, recommended)

The easiest way to run the entire stack (PostgreSQL + backend + frontend) is via:

```sh
make start
```

`make start` is just a shortcut for the more explicit:

```sh
make docker-up
```

or, if you prefer plain Docker commands:

```sh
docker compose up --build -d
```

This command will start:
- **PostgreSQL** database on port `5432`
- **Backend** server on port `8080`
- **Frontend** application on port `3000`

After startup, services will be available at:
- **Frontend UI**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **Swagger documentation**: http://localhost:8080/swagger/index.html

To stop the services launched by `make start`:
```sh
make docker-down
```

or

```sh
docker compose down
```

To view logs:
```sh
docker compose logs -f
```

### Method 2: Hybrid Setup (`make start-local`)

Use this when you want PostgreSQL to run inside Docker but keep the backend and frontend running locally (useful for debugging with hot reload).

```sh
make start-local
```

What happens under the hood:
- `docker compose up postgres -d` starts the database container.
- `.env` (if present) is sourced so that backend/frontend share the same settings.
- Backend runs via `go run ./cmd/server/main.go`.
- Frontend runs via `npm run dev` with `VITE_API_BASE_URL=http://localhost:8080`.
- When you stop the process (Ctrl+C), local processes are terminated and the Postgres container is stopped.

### Method 3: Local Development (without Docker)

If you want to run the service locally without Docker:

#### 2.1. Running Backend

1. Make sure PostgreSQL is running and accessible
2. Create the database:
```sh
createdb calculator
```
or via psql:
```sql
CREATE DATABASE calculator;
```

3. Run migrations (if any):
```sh
psql -d calculator -f migrate/001_init_schema.sql
```

4. Configure environment variables (create a `.env` file):
```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=calculator
```

5. Start the backend:
```sh
go run ./cmd/server/main.go
```

Backend will be available at http://localhost:8080

#### 2.2. Running Frontend

1. Navigate to the frontend directory:
```sh
cd frontend
```

2. Install dependencies:
```sh
npm install
```

3. Start the dev server:
```sh
npm run dev
```

Frontend will be available at http://localhost:5173 (or another port specified by Vite)

#### 2.3. Running Backend and Frontend Together

Use the Makefile command:
```sh
make dev-up
```

This command will start both services simultaneously. To stop:
```sh
make dev-down
```

---

## 4. Check Test Coverage Information  
To check test coverage, run:  
```sh
make test-report
```
This will generate and open a UI report in the browser.

To run all tests:
```sh
make test
```

or

```sh
go test ./...
```

---

## 5. UI  
- **Swagger Button**: Opens Swagger API documentation.  
- **Add Package Sizes**: Use the **X** button to remove a package size and the **✔️** button to submit.  
- **Order Input**: Enter the desired order size in the input field.  
- **Calculate Button**: Click to compute the optimized package distribution.  

---

## 6. API  
The API can be accessed via:  
```
http://localhost:8080/swagger/index.html
```
Alternatively, use the **Swagger Button** in the UI to open the API docs.

### Main endpoints:
- `GET /health` - health check endpoint
- `GET /api/packages` - get list of package sizes
- `POST /api/packages` - add/update package sizes
- `POST /api/calculate` - calculate optimal package distribution

## 7. Deployed service
There is packager deployed publicly here (server side rendered optimised for Render deployment free plaf , source branch is [render-dev](https://github.com/klausborkowski/calculator/tree/render-dev)): [Packager Service](https://calculator-ieo1.onrender.com/app)