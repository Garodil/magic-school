CREATE TABLE professors (
    id INTEGER PRIMARY KEY,
    entity_id INTEGER,
    hiring_date DATE,
    FOREIGN KEY (entity_id) REFERENCES entities(id)
);

CREATE TABLE specialities (
    id INTEGER PRIMARY KEY,
    name TEXT
);

CREATE TABLE professor_specialities (
    professor_id INTEGER,
    speciality_id INTEGER,
    PRIMARY KEY (professor_id, speciality_id),
    FOREIGN KEY (professor_id) REFERENCES professors(id),
    FOREIGN KEY (speciality_id) REFERENCES specialities(id)
);