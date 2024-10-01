# Elabram Backend Recruitment

This is a simple RESTful API written in Go and uses MySQL as database. This API provides endpoints to create, read, update and delete products. Additionally, it uses Redis as caching layer for frequently requested data such as dashboard reports of products.

## Tech Stack

- Go as programming language
- MySQL as database
- Docker as containerization
- Docker Compose as orchestration
- Redis as caching layer

## How to Run

Run `docker-compose up -d` to start the database and Redis in background.
Access the API: Once the containers are running, you can interact with the API.

## Using the API

Base URL
The base URL for the API is http://localhost:8080.

## Endpoints

- `POST /products`: Create a new product
- `GET /products`: Retrieve a paginated list of products
- `GET /products/:id`: Retrieve a product by ID
- `PUT /products/:id`: Update a product by ID
- `DELETE /products/:id`: Delete a product by ID
- `POST /categories`: Create a new product category
- `GET /categories`: Retrieve a list of categories
- `GET /categories/:id`: Retrieve a category by ID
- `GET /reports/products`: Retrieve a report of all products for dashboards

