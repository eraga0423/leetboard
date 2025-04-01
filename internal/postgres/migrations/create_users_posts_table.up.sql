CREATE TABLE "users_posts" (
  "post_id" INT REFERENCES posts(post_id) ON DELETE CASCADE ,
  "user_id" VARCHAR(150)  NOT NULL UNIQUE
);

