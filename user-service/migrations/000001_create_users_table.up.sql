-- migrations/0001_create_users_table.up.sql
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(50) NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password varchar(255) NOT NULl,
                       rating INT DEFAULT 500,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
