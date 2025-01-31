# Author Microservice

## Overview
The **Author** Microservice is part of the **Library Management System**.

## Microservices

- **UserService/AuthService**: Handles user authentication and authorization.
- **BookService**: Manages books and stock.
- **CategoryService**: Manages book categories.
- **AuthorService**: Manages authors.

This microservice is built using:
- **Golang** for the backend.
- **PostgreSQL** for the database.
- **Docker** for containerization and deployment.
- **Docker Hub** for storing Docker images.

---

## **Technologies Used**
- **Programming Language**: Golang.
- **Database**: PostgreSQL.
- **Communication**: gRPC.
- **Middleware**: JWT for authentication.
- **Containerization**: Docker & Docker Compose.

---

## **API Documentation**
### REST API Endpoints
| HTTP Method | Endpoint                        | Description                       |
|-------------|---------------------------------|-----------------------------------|
| `GET`       | `/api/v1/authors`               | Get all authors                   |
| `POST`      | `/api/v1/authors`               | Create a new authors              |
| `GET`       | `/api/v1/authors/:id`           | Get details of a specific authors |
| `PUT`       | `/api/v1/authors/:id`           | Update a specific authors         |
| `DELETE`    | `/api/v1/authors/:id`           | Delete a specific authors         |

---

## Installation

### Prerequisites
- Install [Go](https://go.dev/doc/install)
- Install [PostgreSQL](https://www.postgresql.org/download/)
- Install [Docker](https://docs.docker.com/get-docker/)
- Install [gRPC](https://grpc.io/docs/languages/go/quickstart/)

### Running Without Docker

1. Clone the repository:
   ```sh
   git clone https://github.com/naufalhakm/library-api-author.git
   cd library-api-author
   ```
2. Setup environment variables (.env file):
   ```sh
   DB_HOST=localhost
   DB_PORT=5432
   DB_USERNAME=user
   DB_PASSWORD=password
   DB_DATABASE=library
   ```
3. Run PostgreSQL locally.
4. Start author microservice:
   ```sh
   go run cmd/server/main.go
   ```

### Running With Docker

1. Build and run services:
   ```sh
   docker-compose up -d
   ```

---

### Live Server

The microservice is running at:
http://35.240.139.186:8083/