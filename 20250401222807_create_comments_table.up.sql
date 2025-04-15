
CREATE TABLE IF NOT EXISTS comments (
  comment_id SERIAL PRIMARY KEY,
  post_id INT REFERENCES posts(post_id) ON DELETE CASCADE,
  comment_content TEXT,
  comment_image TEXT,
  comment_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);




