-- Up migration
CREATE TABLE user (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       points INT NOT NULL DEFAULT 0
);