CREATE TABLE IF NOT EXISTS users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(100) UNIQUE
);

ALTER TABLE users ADD UNIQUE (name);