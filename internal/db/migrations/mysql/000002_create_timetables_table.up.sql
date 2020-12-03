CREATE TABLE IF NOT EXISTS Timetables
(
    id            INT          NOT NULL UNIQUE AUTO_INCREMENT,
    name          VARCHAR(255) NOT NULL,
    class_days    INT          NOT NULL DEFAULT 5,
    class_periods INT          NOT NULL DEFAULT 6,
    is_default    BOOLEAN      NOT NULL,
    created_at    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    user_id       VARCHAR(36)  NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users (id),
    PRIMARY KEY (id)
)