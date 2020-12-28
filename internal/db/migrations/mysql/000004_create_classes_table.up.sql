CREATE TABLE IF NOT EXISTS classes
(
    id           INT          NOT NULL UNIQUE AUTO_INCREMENT,
    name         VARCHAR(255) NOT NULL,
    days         INT          NOT NULL,
    periods      INT          NOT NULL,
    style        VARCHAR(64)  NOT NULL,
    room_or_url  VARCHAR(128) NOT NULL,
    teacher      VARCHAR(128) NOT NULL,
    credit       INT          NOT NULL default 0,
    memo         TEXT,
    color        VARCHAR(64)  NOT NULL,
    created_at   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    timetable_id INT          NOT NULL,
    FOREIGN KEY (timetable_id) REFERENCES timetables (id),
    PRIMARY KEY (id)
)