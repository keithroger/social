
CREATE TABLE users(
     user_id INT NOT NULL,
     username VARCHAR (20) NOT NULL,
     name VARCHAR (50),
     email VARCHAR (40) NOT NULL,
     date_created TIMESTAMPTZ,
     PRIMARY KEY (user_id)
);

CREATE TABLE followers (
     follower INT,
     followee INT,
     date_created TIMESTAMPTZ,
     FOREIGN KEY(follower) REFERENCES users(user_id),
     FOREIGN KEY(followee) REFERENCES users(user_id),
     PRIMARY KEY (follower, followee)
);

CREATE TABLE posts (
     user_id INT NOT NULL,
     content VARCHAR(280),
     date_created TIMESTAMPTZ,
     FOREIGN KEY(user_id) REFERENCES users(user_id)
)