CREATE TABLE lessons (
    id INTEGER PRIMARY KEY,
    name TEXT,
    lesson_type_id INTEGER,
    room_id INTEGER,
    FOREIGN KEY (lesson_type_id) REFERENCES lesson_types(id),
    FOREIGN KEY (room_id) REFERENCES rooms(id)
);

CREATE TABLE lesson_types (
    id INTEGER PRIMARY KEY,
    name TEXT,
    allowed_room_type_id INTEGER,
    FOREIGN KEY (allowed_room_type_id) REFERENCES room_types(id)
);

CREATE TABLE lesson_types_room_types (
    lesson_type_id INTEGER,
    room_type_id INTEGER,
    PRIMARY KEY (lesson_type_id, room_type_id),
    FOREIGN KEY (lesson_type_id) REFERENCES lesson_types(id),
    FOREIGN KEY (room_type_id) REFERENCES room_types(id)
);

CREATE TABLE lesson_type_dimensions (
    lesson_type_id INTEGER,
    dimension_id INTEGER,
    PRIMARY KEY (lesson_type_id, dimension_id),
    FOREIGN KEY (lesson_type_id) REFERENCES lesson_types(id),
    FOREIGN KEY (dimension_id) REFERENCES dimensions(id)
);

CREATE TABLE moon_phases (
    id INTEGER PRIMARY KEY,
    name TEXT
);

CREATE TABLE lesson_type_moon_phases (
    lesson_type_id INTEGER,
    moon_phase_id INTEGER,
    PRIMARY KEY (lesson_type_id, moon_phase_id),
    FOREIGN KEY (lesson_type_id) REFERENCES lesson_types(id),
    FOREIGN KEY (moon_phase_id) REFERENCES moon_phases(id)
);

CREATE TABLE seasons (
    id INTEGER PRIMARY KEY,
    name TEXT
);

CREATE TABLE lesson_type_seasons (
    lesson_type_id INTEGER,
    season_id INTEGER,
    PRIMARY KEY (lesson_type_id, season_id),
    FOREIGN KEY (lesson_type_id) REFERENCES lesson_types(id),
    FOREIGN KEY (season_id) REFERENCES seasons(id)
);

CREATE TABLE time_of_day (
    id INTEGER PRIMARY KEY,
    name TEXT
);

CREATE TABLE lesson_type_time_of_day (
    lesson_type_id INTEGER,
    time_of_day_id INTEGER,
    PRIMARY KEY (lesson_type_id, time_of_day_id),
    FOREIGN KEY (lesson_type_id) REFERENCES lesson_types(id),
    FOREIGN KEY (time_of_day_id) REFERENCES time_of_day(id)
);

