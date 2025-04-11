CREATE TABLE IF NOT EXISTS posts (
  post_id SERIAL PRIMARY KEY,
  title VARCHAR(150)  NOT NULL,
  post_content TEXT,
  post_image TEXT,
  post_time TIMESTAMP,
  deletion BOOLEAN DEFAULT FALSE
);
