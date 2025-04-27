CREATE TABLE IF NOT EXISTS users (
  user_id SERIAL PRIMARY KEY,
  name VARCHAR(150)  NOT NULL UNIQUE,
  avatar_url TEXT,
  session_id TEXT
);