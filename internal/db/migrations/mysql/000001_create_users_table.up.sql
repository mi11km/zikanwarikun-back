CREATE TABLE IF NOT EXISTS users
(
    id       VARCHAR(36)  NOT NULL UNIQUE,
    email    VARCHAR(128) NOT NULL UNIQUE,
    password VARCHAR(128) NOT NULL,
    school   VARCHAR(128),
    name     VARCHAR(255),
    PRIMARY KEY (id)
)