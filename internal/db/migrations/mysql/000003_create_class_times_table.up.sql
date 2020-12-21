CREATE TABLE IF NOT EXISTS class_times
(
    id           INT NOT NULL UNIQUE AUTO_INCREMENT,
    periods      INT NOT NULL,
    start_time   TIME,
    end_time     TIME,
    timetable_id INT NOT NULL,
    FOREIGN KEY (timetable_id) REFERENCES timetables (id),
    PRIMARY KEY (id)
)