CREATE DATA IF NOT EXISTS cartdb;
USE cartdb;

CREATE TABLE IF NO EXISTS cart_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT,
    quantity INT
);