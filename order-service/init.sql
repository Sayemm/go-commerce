CREATE DATA IF NOT EXISTS orderdb;
USE orderdb;

CREATE TABLE IF NOT EXISTS orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    created_at DATETIME
);