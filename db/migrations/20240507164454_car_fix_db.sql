-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    email VARCHAR(50) NOT NULL,
    encrypted_password VARCHAR(255) NOT NULL
);
CREATE TABLE servic (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    cost DECIMAL(10, 2) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE services;
DROP TABLE clients;
-- +goose StatementEnd