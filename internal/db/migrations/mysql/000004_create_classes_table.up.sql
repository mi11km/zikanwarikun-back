CREATE TABLE IF NOT EXISTS Classes
(
    id           INT          NOT NULL UNIQUE AUTO_INCREMENT,
    name         VARCHAR(255) NOT NULL,
    day          INT          NOT NULL,
    periods      INT          NOT NULL,
    style        VARCHAR(64)  NOT NULL,
    color        VARCHAR(64)  NOT NULL DEFAULT "orange",
    teacher      VARCHAR(128) NOT NULL,
    credit       INT          NOT NULL default 0,
    room_or_url  VARCHAR(128) NOT NULL,
    memo         TEXT,
    timetable_id INT          NOT NULL,
    FOREIGN KEY (timetable_id) REFERENCES Timetables (id),
    PRIMARY KEY (id)
)