
DELETE FROM users;
ALTER SEQUENCE users_user_id_seq RESTART;
ALTER SEQUENCE posts_post_id_seq RESTART;
ALTER SEQUENCE comments_comment_id_seq RESTART;
-- Users
INSERT INTO
    users(username, profile_name)
VALUES
    ('johnathan123', 'jonathan johnson'),
    ('alice3', 'alice alexis'),
    ('smith', 'Elizabeth Smith'),
    ('molly1234', 'Molly Monday'),
    ('123sarah', 'Sarrah Samson'),
    ('321Josh1', 'Joshua Jackobs'),
    ('1Jjake', 'Jake Jackson'),
    ('theChris', 'Chistopher Charlie'),
    ('elCharlie', 'Charlie Carson'),
    ('itsDanny', 'Daniella Daniels'),
    ('bran', 'Brandon Branson'),
    ('brandy', 'Brandy Bronson'),
    ('robinson', 'Robin Robinson'),
    ('Ashley', 'Ashley Ace'),
    ('itsRach', 'Rachel Randomx'),
    ('somebody', 'Joe Jameson'),
    ('aPerson', 'Jen Jetson'),
    ('laurenL', 'Lauren Long'),
    ('jazzy', 'Jasmine Le'),
    ('Hannah', 'Hannah Nguyen'),
    ('joseRod', 'Jose Rodriquez');

-- Followers
INSERT INTO
    followers(follower_id, followee_id)
VALUES
    (1, 3),
    (1, 4),
    (2, 5),
    (2, 6),
    (2, 7),
    (2, 8),
    (3, 1),
    (3, 2),
    (3, 4),
    (4, 1),
    (4, 2),
    (4, 3),
    (4, 5),
    (4, 6),
    (4, 7),
    (4, 8),
    (4, 9),
    (4, 10),
    (4, 11),
    (4, 12),
    (4, 13);


INSERT INTO posts(user_id, body, likes)
VALUES
(1, 'Some body content', 0),
(1, 'more content by user 1', 0),
(1, 'more', 0),
(1, 'more', 0),
(2, 'Some more posting', 0),
(3, 'A third post', 0),
(4, 'Fourth post',  0),
(5, 'Fifth post', 0),
(6, 'Sixth post', 0),
(7, 'Seventh post', 0);

INSERT INTO comments(post_id, user_id, body)
 VALUES
 (2, 4, 'first comment'),
 (2, 4, 'second comment'),
 (2, 4, 'third comment'),
 (2, 4, 'fourth comment'),
 (2, 4, 'fifth comment'),
 (2, 4, 'sixth comment'),
 (2, 4, 'seventh comment'),
 (2, 4, 'eighth comment');