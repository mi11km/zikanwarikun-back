CREATE TABLE IF NOT EXISTS user_timetables
(
    id           INT         NOT NULL UNIQUE AUTO_INCREMENT,
    user_id      VARCHAR(36) NOT NULL,
    timetable_id INT         NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (timetable_id) REFERENCES timetables (id),
    PRIMARY KEY (id)
)