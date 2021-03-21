CREATE TABLE `users` (
    id BINARY(16) NOT NULL UNIQUE DEFAULT UNHEX(REPLACE(UUID(), '-', '')),
    username VARCHAR(32) NOT NULL UNIQUE,
    email VARCHAR(320) NOT NULL UNIQUE,
    password VARCHAR(60) NOT NULL,
    level_id INT NOT NULL DEFAULT 0,
    last_seen DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    website VARCHAR(2048),
    biography TEXT,
    views INT NOT NULL DEFAULT 0,
    uploads INT NOT NULL DEFAULT 0,
    premium BOOL NOT NULL DEFAULT false,
    md_at_home BOOL NOT NULL DEFAULT false,
    avatar_url VARCHAR(2048) NOT NULL,

    joined_at DATETIME DEFAULT NOW(),
    update_at DATETIME DEFAULT NOW() ON UPDATE NOW()
);

INSERT INTO `users` (username, email, password, website, biography, avatar_url)
VALUES (
    'root',
    'root@example.com',
    '$2a$10$EMCB89WuPHdGMMY.3Xy2yuxHsjVW36sDuq01y4lyU2zcH1JHJSen6',
    'https://example.com',
    'The root account used on initial installation.',
    'https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50.jpg'
);
