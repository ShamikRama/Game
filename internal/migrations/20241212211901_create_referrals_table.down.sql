CREATE TABLE referrals (
                           id SERIAL PRIMARY KEY,
                           user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE, -- Пользователь, который ввел код
                           referrer_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE, -- Пользователь, который предоставил код
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);