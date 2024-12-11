CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       user_id INT NOT NULL,
                       type VARCHAR(255) NOT NULL,
                       exp INT NOT NULL,
                       complete BOOLEAN NOT NULL DEFAULT FALSE,
                       CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);