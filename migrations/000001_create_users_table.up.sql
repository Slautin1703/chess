-- migrations/0001_create_users_table.up.sql
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(50) NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL,
                       password varchar(100) NOT NULl,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
