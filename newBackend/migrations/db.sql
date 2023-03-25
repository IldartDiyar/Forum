CREATE TABLE IF NOT EXISTS user(
		userId  INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		password text, 	
		email text UNIQUE
		-- token TEXT ,
		-- expiresAt DATETIME
);
CREATE TABLE IF NOT EXISTS user_avatar(
	avatarId INTEGER PRIMARY KEY AUTOINCREMENT,
	userId INTEGER UNIQUE,
	base  TEXT,
	FOREIGN KEY (userId) REFERENCES user(userId) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_sessions(
  token TEXT PRIMARY KEY,
  expiresAt TEXT,
  userId INTEGER,
  FOREIGN KEY (userId) REFERENCES user(userId)
);
CREATE TABLE IF NOT EXISTS posts(
		postId INTEGER PRIMARY KEY AUTOINCREMENT,
		author REFERENCES user(username),
		title text UNIQUE,
		content text,
		like INTEGER DEFAULT 0,
		dislike INTEGER DEFAULT 0,
		creationDate TEXT,
		category_id INTEGER,
		ImageName TEXT DEFAULT NULL,
		ImageBase TEXT DEFAULT NULL,
		FOREIGN KEY (category_id) REFERENCES posts_category(category_id)
);
CREATE TABLE IF NOT EXISTS posts_category(
    	category_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		category_name TEXT UNIQUE
);

INSERT INTO posts_category (category_name)
SELECT 'Alem'
WHERE NOT EXISTS (SELECT * FROM posts_category WHERE category_name = 'Alem');
 
INSERT INTO posts_category (category_name)
SELECT 'Golang'
WHERE NOT EXISTS (SELECT * FROM posts_category WHERE category_name = 'Golang');
 
INSERT INTO posts_category (category_name)
SELECT 'JS'
WHERE NOT EXISTS (SELECT * FROM posts_category WHERE category_name = 'JS');
 
INSERT INTO posts_category (category_name)
SELECT 'Rust'
WHERE NOT EXISTS (SELECT * FROM posts_category WHERE category_name = 'Rust');


CREATE TABLE IF NOT EXISTS comments(
    	commentsId INTEGER PRIMARY KEY AUTOINCREMENT,
    	postId INTEGER,
    	author TEXT,
    	content TEXT,
		like INTEGER DEFAULT 0,
		dislike INTEGER DEFAULT 0,
		creationDate TEXT DEFAULT NULL,
    	FOREIGN KEY (postId)  REFERENCES posts(postId) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS likesPost(
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId INTEGER,
    	postId INTEGER,
		like1 INT,
    	FOREIGN KEY (postId) REFERENCES posts(postId) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS likesComment(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
    	userId INTEGER,
    	commentsId INTEGER,
		like1 INT,
    	FOREIGN KEY (commentsId) REFERENCES comments(commentsId) ON DELETE CASCADE
);