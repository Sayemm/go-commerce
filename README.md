# 🛒 Go-Commerce — Product Service

This is the **Product Service** of the `go-commerce` project — a modular e-commerce system built using Go, MySQL, Docker, and Docker Compose. This microservice handles product-related operations like listing and adding products.

> ⚠️ This is **not a complete application**, but the **first service** in a growing microservices architecture. Future services like Cart, Order, Payment, and Authentication will be added.

---

## ✅ Current Features

- RESTful API using Gin
- `GET /products` — List all products
- `POST /products` — Add a new product
- MySQL integration via SQLx
- Dockerized using Docker & Docker Compose

---

## 🗂️ Project Structure

```
go-commerce/
├── backend/
│   ├── db/              # DB connection logic
│   ├── handlers/        # HTTP handlers (controllers)
│   ├── models/          # Data models (e.g., Product)
│   ├── routes/          # API routes
│   ├── main.go          # Entry point
│   ├── go.mod
│   └── go.sum
├── Dockerfile           # Build config for product service
├── docker-compose.yml   # Defines product-service + MySQL setup
└── README.md            # Project overview (this file)
```

---

## 🔗 API Endpoints

### ➕ Add Product

```http
POST /products
Content-Type: application/json

{
  "name": "Laptop",
  "price": 999.99,
  "quantity": 10
}
```

### 📦 Get Products

```http
GET /products
```

Returns an array of products.

---

## 🐳 Docker & Compose Commands

### ⚙️ Build & Run Services

```bash
docker-compose up --build
```

### ▶️ Start Existing Containers

```bash
docker-compose start
```

### ⏹️ Stop Running Containers

```bash
docker-compose stop
```

### 🔁 Restart Containers

```bash
docker-compose restart
```

### 📄 View Logs

```bash
docker logs go-commerce-product-service-1
```

---

## 🧠 What I've Learned So Far

### 🔹 Go (Golang)
- Structs, modules, packages
- SQLx for DB interaction
- Clean folder structure for scalability

### 🔹 Gin Web Framework
- Routing
- JSON parsing & response
- Middleware usage (Logger, Recovery)

### 🔹 Docker
- Creating Dockerfiles
- Containerizing Go apps
- Connecting services via Docker Compose

---

## 🛣️ Roadmap

- [ ] Add Cart Service
- [ ] Add Order Service
- [ ] Add Authentication (JWT)
- [ ] Add Payment Gateway Integration
- [ ] Add React Frontend
- [ ] CI/CD + Deployment via Kubernetes

---

## ✍️ Author

**Sayem** — Aspiring backend engineer learning Go by building real-world projects.
