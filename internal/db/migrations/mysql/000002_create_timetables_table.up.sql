CREATE TABLE IF NOT EXISTS timetables
(
    id            INT          NOT NULL UNIQUE AUTO_INCREMENT,
    name          VARCHAR(255) NOT NULL,
    class_days    INT          NOT NULL,
    class_periods INT          NOT NULL,
    is_default    BOOLEAN      NOT NULL,
    created_at    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
)