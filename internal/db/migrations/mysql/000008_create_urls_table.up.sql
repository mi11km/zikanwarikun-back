CREATE TABLE IF NOT EXISTS urls
(
    id         INT          NOT NULL UNIQUE AUTO_INCREMENT,
    name       VARCHAR(255) NOT NULL,
    url        VARCHAR(255) NOT NULL,
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    class_id   INT          NOT NULL,
    FOREIGN KEY (class_id) REFERENCES classes (id),
    PRIMARY KEY (id)
)