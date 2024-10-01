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


### Complex MySQL Query Optimization

1. Create a Relational Schema
  tables: products, categories, orders, customers, and order_items.

```sql
-- Create products table
CREATE TABLE products (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    description TEXT,
    price DECIMAL(10, 2),
    category_id INT,
    stock_quantity INT,
    is_active BOOLEAN DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- Create categories table
CREATE TABLE categories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create customers table
CREATE TABLE customers (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    email VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create orders table
CREATE TABLE orders (
    id INT PRIMARY KEY AUTO_INCREMENT,
    customer_id INT,
    total_price DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);

-- Create order_items table
CREATE TABLE order_items (
    id INT PRIMARY KEY AUTO_INCREMENT,
    order_id INT,
    product_id INT,
    quantity INT,
    unit_price DECIMAL(10, 2),
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);
```


2. Writing Optimized Queries
    1. Retrieve All Products with Their Category Details and Total Sold Quantities
This query retrieves product details, including category information and the total quantity sold from order_items.

```sql
SELECT p.id, p.name, p.description, p.price, p.stock_quantity, c.name AS category_name,
       COALESCE(SUM(oi.quantity), 0) AS total_sold_quantity
FROM products p
LEFT JOIN categories c ON p.category_id = c.id
LEFT JOIN order_items oi ON p.id = oi.product_id
GROUP BY p.id;
```

#### Optimizations:
Indexes:
Add an index on order_items(product_id) for faster joins between products and order_items.
Add an index on products(category_id) for faster category lookups.

```sql
CREATE INDEX idx_product_id ON order_items(product_id);
CREATE INDEX idx_category_id ON products(category_id);
```

2. Generate a Report of the Top 10 Customers Based on Total Amount Spent Across All Orders
This query retrieves the top 10 customers who have spent the most across all their orders.

```sql
SELECT c.id, c.name, c.email, COALESCE(SUM(o.total_price), 0) AS total_spent
FROM customers c
LEFT JOIN orders o ON c.id = o.customer_id
GROUP BY c.id
ORDER BY total_spent DESC
LIMIT 10;
```

#### Optimizations:
Indexes:
Add an index on orders(customer_id) for faster lookups when joining customers and orders.

```sql
CREATE INDEX idx_customer_id ON orders(customer_id);
```

3. Fetching Order History with Related Products and Customers Using Efficient Indexing and Joins
This query fetches order history with details of related products and customers, showing what each customer ordered and the associated products.

```sql
SELECT o.id AS order_id, c.name AS customer_name, c.email, 
       o.total_price, o.created_at AS order_date,
       p.name AS product_name, oi.quantity, oi.unit_price
FROM orders o
JOIN customers c ON o.customer_id = c.id
JOIN order_items oi ON o.id = oi.order_id
JOIN products p ON oi.product_id = p.id
ORDER BY o.created_at DESC;
```

#### Optimizations:
Indexes:
Add indexes on order_items(order_id) and order_items(product_id) for faster joins between orders and order_items, and order_items with products.
Add an index on orders(customer_id) for efficient lookups when fetching customer order history.

```sql
CREATE INDEX idx_order_id ON order_items(order_id);
CREATE INDEX idx_product_id ON order_items(product_id);
CREATE INDEX idx_customer_id ON orders(customer_id);
```

#### Additional Performance Considerations
1. Indexes:
Ensure all foreign keys are indexed. This includes:

orders(customer_id)
order_items(order_id)
order_items(product_id)
products(category_id)
Indexing these fields ensures faster join operations.

2. Partitioning:
Partition large tables like orders by date (e.g., yearly or monthly) if they grow significantly over time. This helps speed up queries involving time ranges (e.g., order history).

```sql
ALTER TABLE orders PARTITION BY RANGE (YEAR(created_at)) (
  PARTITION p2020 VALUES LESS THAN (2021),
  PARTITION p2021 VALUES LESS THAN (2022),
  PARTITION p2022 VALUES LESS THAN (2023)
);
```

3. Materialized Views:
Use materialized views for frequently accessed reports, such as customer spending or product sales statistics. These views can store precomputed results and reduce the need for complex queries in real-time.

Example:

```sql
CREATE MATERIALIZED VIEW customer_spending AS
SELECT customer_id, SUM(total_price) AS total_spent
FROM orders
GROUP BY customer_id;
```

4. Avoid N+1 Problems:
Use joins instead of fetching data in multiple queries.
Avoid fetching related data in loops or in multiple smaller queries, which can degrade performance (N+1 problem).

5. Query Execution Plan:
Always check the query’s execution plan using EXPLAIN to understand how a query will be executed by the database. It provides insights into the query execution plan, showing how the database retrieves data, which can help identify performance bottlenecks and optimize queries.

```sql
EXPLAIN SELECT * FROM orders WHERE customer_id = 1;
```