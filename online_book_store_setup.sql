-- Drop the database if it exists
DROP DATABASE IF EXISTS online_bookstore;

-- Create the database
CREATE DATABASE online_bookstore;

-- Use the database
USE online_bookstore;

-- Create Authors Table
CREATE TABLE authors (
    author_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    birthdate DATE
);

-- Create Books Table
CREATE TABLE books (
    book_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    author_id INT,
    genre VARCHAR(100),
    published_date DATE,
    price DECIMAL(10, 2),
    FOREIGN KEY (author_id) REFERENCES authors(author_id)
);

-- Create Customers Table
CREATE TABLE customers (
    customer_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    phone VARCHAR(15),
    address VARCHAR(255)
);

-- Create Orders Table
CREATE TABLE orders (
    order_id INT AUTO_INCREMENT PRIMARY KEY,
    customer_id INT,
    order_date DATE,
    total_amount DECIMAL(10, 2),
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);

-- Create OrderItems Table
CREATE TABLE order_items (
    order_item_id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT,
    book_id INT,
    quantity INT,
    price DECIMAL(10, 2),
    FOREIGN KEY (order_id) REFERENCES orders(order_id),
    FOREIGN KEY (book_id) REFERENCES books(book_id)
);

-- Insert Data into Authors
INSERT INTO authors (name, birthdate) VALUES
('J.K. Rowling', '1965-07-31'),
('George R.R. Martin', '1948-09-20'),
('J.R.R. Tolkien', '1892-01-03');

-- Insert Data into Books
INSERT INTO books (title, author_id, genre, published_date, price) VALUES
('Harry Potter and the Sorcerer\'s Stone', 1, 'Fantasy', '1997-06-26', 19.99),
('A Game of Thrones', 2, 'Fantasy', '1996-08-06', 22.99),
('The Hobbit', 3, 'Fantasy', '1937-09-21', 14.99);

-- Insert Data into Customers
INSERT INTO customers (name, email, phone, address) VALUES
('John Doe', 'john.doe@example.com', '123-456-7890', '123 Maple St'),
('Jane Smith', 'jane.smith@example.com', '987-654-3210', '456 Oak St');

-- Insert Data into Orders
INSERT INTO orders (customer_id, order_date, total_amount) VALUES
(1, '2024-05-20', 37.98),
(2, '2024-05-21', 14.99);

-- Insert Data into OrderItems
INSERT INTO order_items (order_id, book_id, quantity, price) VALUES
(1, 1, 1, 19.99),
(1, 3, 1, 14.99),
(2, 3, 1, 14.99);

-- Create a stored procedure to create index if not exists
DELIMITER //
CREATE PROCEDURE CreateIndexIfNotExists(IN index_name VARCHAR(255), IN table_name VARCHAR(255), IN column_name VARCHAR(255))
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.statistics
        WHERE table_schema = DATABASE() AND
              table_name = table_name AND
              index_name = index_name
    ) THEN
        SET @sql = CONCAT('CREATE INDEX ', index_name, ' ON ', table_name, '(', column_name, ')');
        PREPARE stmt FROM @sql;
        EXECUTE stmt;
        DEALLOCATE PREPARE stmt;
    END IF;
END //
DELIMITER ;

-- Call the stored procedure to create indexes if they do not exist
CALL CreateIndexIfNotExists('idx_books_title', 'books', 'title');
CALL CreateIndexIfNotExists('idx_customers_email', 'customers', 'email');
