CREATE TABLE IF NOT EXISTS users (
  user_id SERIAL PRIMARY KEY,
  name VARCHAR(150)  NOT NULL UNIQUE,
  user_avatar TEXT,
  metadata JSONB
);