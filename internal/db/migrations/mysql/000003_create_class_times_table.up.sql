CREATE TABLE IF NOT EXISTS class_times
(
    id           INT      NOT NULL UNIQUE AUTO_INCREMENT,
    periods      INT      NOT NULL,
    start_time   TIME     NOT NULL,
    end_time     TIME     NOT NULL,
    created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   DATETIME,
    timetable_id INT      NOT NULL,
    FOREIGN KEY (timetable_id) REFERENCES timetables (id),
    PRIMARY KEY (id)
)