CREATE TABLE IF NOT EXISTS users_posts (
  post_id INT REFERENCES posts(post_id) ON DELETE CASCADE,
  user_id INT NOT NULL CHECK (user_id  > 0)
);

