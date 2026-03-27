-- +goose Up
CREATE TABLE users (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_name VARCHAR(20) NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE users;