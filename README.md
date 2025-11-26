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

To run the service, you need to install:
- **Docker** and **Docker Compose** (for Docker-based setup)
- **Go 1.22+** (for local backend development)
- **Node.js 18+** and **npm** (for local frontend development)
- **PostgreSQL 16+** (for local setup, or use Docker)

---

## 3. Running the Service

### Method 1: Docker Setup (Recommended)

The easiest way to run the entire service is using Docker Compose:

```sh
make docker-up
```

or

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

To stop the services:
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

### Method 2: Local Development (without Docker)

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
There is packager deployed publicly here: [Packager Service](https://packager-0e6j.onrender.com/app)