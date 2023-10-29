-- CREATE DATABASE mobile_phone_reviews;

USE mobile_phone_reviews;

CREATE TABLE brands (
    brand_id INT PRIMARY KEY AUTO_INCREMENT,
    brand_name VARCHAR(255) NOT NULL
);

CREATE TABLE phones (
    phone_id INT PRIMARY KEY AUTO_INCREMENT,
    brand_id INT,
    phone_model VARCHAR(255) NOT NULL,
    phone_category VARCHAR(50) NOT NULL,
    FOREIGN KEY (brand_id) REFERENCES brands(brand_id)
);

CREATE TABLE customers (
    customer_id INT PRIMARY KEY AUTO_INCREMENT,
    customer_name VARCHAR(255) NOT NULL
);

CREATE TABLE reviews (
    review_id INT PRIMARY KEY AUTO_INCREMENT,
    customer_id INT,
    phone_id INT,
    review_text TEXT,
    review_rating DECIMAL(3, 2),
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id),
    FOREIGN KEY (phone_id) REFERENCES phones(phone_id)
);

CREATE TABLE phone_specifications (
    specification_id INT PRIMARY KEY AUTO_INCREMENT,
    phone_id INT,
    specification_name VARCHAR(255),
    specification_value TEXT,
    FOREIGN KEY (phone_id) REFERENCES phones(phone_id)
);
