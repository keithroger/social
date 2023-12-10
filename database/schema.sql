CREATE TABLE users(
     user_id BIGSERIAL,
     username VARCHAR (20) NOT NULL,
     profile_name VARCHAR (50),
     date_created TIMESTAMPTZ,
     PRIMARY KEY (user_id)
);

CREATE TABLE followers (
     follower BIGSERIAL,
     followee BIGSERIAL,
     date_created TIMESTAMPTZ,
     FOREIGN KEY(follower) REFERENCES users(user_id),
     FOREIGN KEY(followee) REFERENCES users(user_id),
     PRIMARY KEY (follower, followee)
);

CREATE TABLE posts (
     post_id BIGSERIAL,
     user_id BIGSERIAL,
     msg VARCHAR(300),
     likes INT,
     date_created TIMESTAMPTZ,
     FOREIGN KEY(user_id) REFERENCES users(user_id),
     PRIMARY KEY (post_id)
);

CREATE TABLE likes (
     post_id BIGSERIAL,
     user_id BIGSERIAL,
     FOREIGN KEY(post_id) REFERENCES posts(post_id),
     FOREIGN KEY(user_id) REFERENCES users(user_id),
     PRIMARY KEY(post_id, user_id)
);

CREATE TABLE comments (
     comment_id BIGSERIAL,
     post_id BIGSERIAL,
     user_id INT NOT NULL,
     msg VARCHAR(300),
     date_created TIMESTAMPTZ,
     FOREIGN KEY(post_id) REFERENCES posts(post_id),
     FOREIGN KEY(user_id) REFERENCES users(user_id)
);