CREATE TABLE IF NOT EXISTS avatars (
  character_id INT UNIQUE,
  character_name TEXT,
  character_image TEXT,
  status BOOLEAN DEFAULT FALSE
);