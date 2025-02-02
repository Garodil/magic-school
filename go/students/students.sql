CREATE TABLE students (
    id INTEGER PRIMARY KEY,
    entity_id INTEGER,
    enrollment_date DATE,
    FOREIGN KEY (entity_id) REFERENCES entities(id)
);