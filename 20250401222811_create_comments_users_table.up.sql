CREATE TABLE IF NOT EXISTS comments_users (
  comment_id INT REFERENCES comments(comment_id) ON DELETE CASCADE,
  user_id INT NOT NULL CHECK (user_id  > 0)
);

