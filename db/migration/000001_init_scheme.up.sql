CREATE TABLE users (
  id bigserial PRIMARY KEY,
  email VARCHAR(100) UNIQUE NOT NULL,
  full_name VARCHAR(50),
  hashed_password VARCHAR(100) NOT NULL,
  city VARCHAR(50),
  state VARCHAR(50),
  country VARCHAR(50),
  image_url varchar(500),
  created_at timestamptz NOT NULL DEFAULT (now()),
  updated_at TIMESTAMP
);

CREATE TABLE pets (
    id bigserial PRIMARY KEY,
    name varchar(50) NOT NULL,
    about text,
    user_id bigint REFERENCES users(id) NOT NULL,
    age integer NOT NULL,
    gender varchar(20) NOT NULL,
    pet_type varchar(50) NOT NULL,
    breed varchar(50),
    image_url varchar(500),
    medical_condition varchar(50),
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at TIMESTAMP
);

CREATE TABLE posts (
    id bigserial PRIMARY KEY,
    title varchar(100),
    body text,
    user_id bigint REFERENCES users(id) NOT NULL,
    image_url varchar(500),
    status varchar(20),
    likes_count integer NOT NULL DEFAULT 0 CHECK (likes_count >= 0),
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at TIMESTAMP
);

CREATE TABLE likes (
    id bigserial PRIMARY KEY,
    liked_post_id bigint REFERENCES posts(id) NOT NULL,
    user_id bigint REFERENCES users(id) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT unique_user_post_likes UNIQUE (user_id, liked_post_id)
);

CREATE INDEX idx_user_id_pets ON pets(user_id);
CREATE INDEX idx_user_id_posts ON posts(user_id);
CREATE INDEX idx_user_id_likes ON likes(user_id);
CREATE INDEX idx_liked_post_id_likes ON likes(liked_post_id);

CREATE OR REPLACE FUNCTION decrement_likes_count() RETURNS TRIGGER AS $$
BEGIN
  UPDATE posts SET likes_count = likes_count - 1 WHERE id = OLD.liked_post_id;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER decrement_likes_count
AFTER DELETE ON likes
FOR EACH ROW
EXECUTE FUNCTION decrement_likes_count();

CREATE OR REPLACE FUNCTION increment_likes_count() RETURNS TRIGGER AS $$
BEGIN
  UPDATE posts SET likes_count = likes_count + 1 WHERE id = NEW.liked_post_id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER increment_likes_count
AFTER INSERT ON likes
FOR EACH ROW
EXECUTE FUNCTION increment_likes_count();
