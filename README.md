# 🛒 E-commerce Checkout System – Microservices in Go

This is a containerized **E-commerce Checkout System** built using **Go (Golang)**. It follows a **microservices architecture** with independent REST APIs for **Product**, **Cart**, and **Order** services. Each service has its own **MySQL** database and is orchestrated using **Docker Compose**.

---

## 🚀 Project Overview

The system consists of three microservices:

- **Product Service**: Manage product creation and listing.
- **Cart Service**: Handle cart operations.
- **Order Service**: Place and view orders.

Each microservice is:

- Written in Go using the **Gin** framework
- Uses **sqlx** for database interactions
- Runs in its own Docker container

---

## 🛠 Tech Stack

- **Go (Golang)**
- **Gin** (Web Framework)
- **MySQL**
- **sqlx** (Database library)
- **Docker & Docker Compose**

---

## ✅ Features

### 🛍 Product Service
- Add new products
- List all products

### 🧺 Cart Service
- Add products to cart
- View cart items

### 📦 Order Service
- Place orders
- View order history

---

## ⚙️ Getting Started

### Steps to Run

1. Clone the Repository

```bash
git clone https://github.com/Sayemm/go-commerce.git
cd go-commerce
```

2. Run the Project

```bash
docker-compose up --build
```

3. Access Services

Product Service: http://localhost:8000  
Cart Service: http://localhost:8001  
Order Service: http://localhost:8002

## Learnings

- Building REST APIs in Go using Gin
- Using sqlx to interact with MySQL databases
- Writing modular, scalable microservices
- Dockerizing Go apps
- Managing multi-container apps using Docker Compose

## Upcoming Features

- Product inventory deduction on order placement
- Cart → Order integration with item validation
- gRPC or Kafka for inter-service communication
- User authentication service with JWT
- React frontend for full-stack experience
