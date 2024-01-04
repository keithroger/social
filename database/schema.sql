CREATE TABLE users(
     user_id BIGSERIAL,
     username VARCHAR (20) NOT NULL,
     profile_name VARCHAR (50),
     date_created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
     PRIMARY KEY (user_id)
);

CREATE TABLE followers (
     follower_id BIGSERIAL,
     followee_id BIGSERIAL,
     date_created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
     FOREIGN KEY(follower_id) REFERENCES users(user_id) ON DELETE CASCADE,
     FOREIGN KEY(followee_id) REFERENCES users(user_id) ON DELETE CASCADE,
     PRIMARY KEY (follower_id, followee_id)
);

CREATE TABLE posts (
     post_id BIGSERIAL,
     user_id BIGSERIAL,
     body VARCHAR(300),
     likes INT,
     date_created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
     FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE,
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
     body VARCHAR(300),
     date_created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
     FOREIGN KEY(post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
     FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE
);