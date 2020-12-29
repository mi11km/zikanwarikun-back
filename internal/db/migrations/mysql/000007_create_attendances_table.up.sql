CREATE TABLE IF NOT EXISTS attendances
(
    id         INT      NOT NULL UNIQUE AUTO_INCREMENT,
    attend     INT      NOT NULL DEFAULT 0,
    absent     INT      NOT NULL DEFAULT 0,
    late       INT      NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    class_id   INT      NOT NULL,
    FOREIGN KEY (class_id) REFERENCES classes (id),
    PRIMARY KEY (id)
)