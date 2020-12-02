CREATE TABLE IF NOT EXISTS Users
(
    id       INT          NOT NULL UNIQUE AUTO_INCREMENT,
    email    VARCHAR(128) NOT NULL UNIQUE,
    password VARCHAR(128) NOT NULL,
    school   VARCHAR(128),
    name     VARCHAR(255),
    PRIMARY KEY (id)
)