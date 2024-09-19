INSERT INTO user (name, nick, email, password) VALUES
('user1', 'user1', 'user1@gmail.com', '$2a$10$3lnNPbcX.xm0vBD.WO8acevC9.sdtdMsEI4WQdAh4Vi5R9M23lCVy'),
('user2', 'user2', 'user2@gmail.com', '$2a$10$3lnNPbcX.xm0vBD.WO8acevC9.sdtdMsEI4WQdAh4Vi5R9M23lCVy'),
('user3', 'user3', 'user3@gmail.com', '$2a$10$3lnNPbcX.xm0vBD.WO8acevC9.sdtdMsEI4WQdAh4Vi5R9M23lCVy'),
('user4', 'user4', 'user4@gmail.com', '$2a$10$3lnNPbcX.xm0vBD.WO8acevC9.sdtdMsEI4WQdAh4Vi5R9M23lCVy');
 
INSERT INTO follower (user_id, follower_id) VALUES
(1, 4),
(2, 3);

INSERT INTO post (author_id, title, description) VALUES
(1, 'Post 1', 'content 1'),
(2, 'Post 2', 'content 2'),
(3, 'Post 3', 'content 3'),
(4, 'Post 4', 'content 4');