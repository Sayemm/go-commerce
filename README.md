# E-commerce Checkout System - Microservices in Golang

This project is a containerized E-commerce Checkout System built using Go (Golang) with REST APIs for the Product, Cart, and Order services. Each service is independent, uses its own MySQL database, and is managed using Docker Compose.

## Project Overview

The system consists of three microservices:

- Product Service: Handles product creation and listing.
- Cart Service: Manages items added to a user's cart.
- Order Service: Handles order placement and viewing.

Each service is written in Go using the gin framework, uses sqlx for database operations, and has its own Docker container.

## Folder Structure

go-commerce/
├── product-service/
│   ├── main.go
│   ├── db/
│   ├── models/
│   ├── handlers/
│   ├── routes/
│   └── Dockerfile
├── cart-service/
│   ├── main.go
│   ├── db/
│   ├── models/
│   ├── handlers/
│   ├── routes/
│   └── Dockerfile
├── order-service/
│   ├── main.go
│   ├── db/
│   ├── models/
│   ├── handlers/
│   ├── routes/
│   └── Dockerfile
├── docker-compose.yml
└── README.md

## Tech Stack

- Go (Golang)
- Gin
- MySQL
- sqlx
- Docker & Docker Compose

## Features Implemented

Product Service:
- Add new products
- List all products

Cart Service:
- Add products to cart
- List items in cart

Order Service:
- Place orders
- List order history

## Setup Instructions

### Steps to Run

1. Clone the Repository

git clone https://github.com/Sayemm/go-commerce.git
cd go-commerce

2. Run the Project

docker-compose up --build

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
