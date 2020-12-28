CREATE TABLE IF NOT EXISTS todos
(
    id          INT         NOT NULL UNIQUE AUTO_INCREMENT,
    kind        VARCHAR(64) NOT NULL,
    deadline    DATETIME    NOT NULL,
    is_done     BOOLEAN     NOT NULL DEFAULT FALSE,
    memo        TEXT,
    is_repeated BOOLEAN     NOT NULL DEFAULT false,
    created_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    class_id    INT         NOT NULL,
    FOREIGN KEY (class_id) REFERENCES classes (id),
    PRIMARY KEY (id)
)