CREATE TABLE IF NOT EXISTS comments_users (
  comment_id INT REFERENCES comments(comment_id) ON DELETE CASCADE ,
  user_id VARCHAR(150)  NOT NULL UNIQUE
);

