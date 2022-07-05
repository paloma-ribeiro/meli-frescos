-- Active: 1657031674704@@127.0.0.1@3306@meli_frescos
DROP SCHEMA IF EXISTS meli_frescos;
CREATE SCHEMA meli_frescos;
USE meli_frescos;

CREATE TABLE `employees` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `card_number_id` VARCHAR(255) NOT NULL UNIQUE,
  `first_name` VARCHAR(255) NOT NULL,
  `last_name` VARCHAR(255) NOT NULL,
  `warehouse_id` int NOT NULL
);