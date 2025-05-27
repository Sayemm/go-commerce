CREATE DATABASE IF NOT EXISTS cartdb;
USE cartdb;

CREATE TABLE IF NOT EXISTS cart_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT,
    quantity INT
);