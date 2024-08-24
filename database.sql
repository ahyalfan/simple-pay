-- Active: 1723868453146@@127.0.0.1@5432@godof
CREATE DATABASE godof;

CREATE Table users (
    id SERIAL NOT NULL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    -- email_verified_at timestamp(6) without time zone
);

ALTER Table users
ADD COLUMN email_verified_at TIMESTAMP(0) without time zone;

ALTER Table users ADD COLUMN email VARCHAR(255) UNIQUE NOT NULL;