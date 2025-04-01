CREATE TABLE ID NOT EXISTS subcomments (
  comment_parent INT REFERENCES comments(comment_id) ON DELETE CASCADE ,
  comment_child INT REFERENCES comments(comment_id) ON DELETE CASCADE 
);