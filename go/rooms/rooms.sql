CREATE TABLE rooms ( -- Таблица кабинетов
    id INTEGER PRIMARY KEY,
    name TEXT
);

CREATE TABLE dimensions ( -- Таблица измерений
    id INTEGER PRIMARY KEY,
    name TEXT,
    full_name TEXT,
    sky_color TEXT,
    temperature INTEGER,
    magicality INTEGER,
    danger_level INTEGER,
    description TEXT,
    danger_description TEXT
);

CREATE TABLE room_dimensions ( -- Таблица, связующая кабинеты и измерения
    room_id INTEGER,
    dimension_id INTEGER,
    PRIMARY KEY(room_id, dimension_id),
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (dimension_id) REFERENCES dimensions(id)
);

CREATE TABLE floors ( -- Таблица этажей
    id INTEGER PRIMARY KEY,
    name TEXT,
    full_name TEXT,
    temperature INTEGER,
    magicality INTEGER,
    danger_level INTEGER,
    description TEXT,
    danger_description TEXT,
    corridor_room_id INTEGER,
    age INTEGER,
    created_at DATE,
    FOREIGN KEY (corridor_room_id) REFERENCES rooms(id)
);

CREATE TABLE room_floors ( -- Связующая таблица кабинетов и этажей
    room_id INTEGER,
    floor_id INTEGER,
    PRIMARY KEY(room_id, floor_id),
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (floor_id) REFERENCES floors(id)
);

CREATE TABLE room_types ( -- Таблица типов кабинетов
    id INTEGER PRIMARY KEY,
    name TEXT
);

CREATE TABLE room_room_types ( -- Таблица кабинетов и типов кабинетов
    room_id INTEGER,
    room_type_id INTEGER,
    PRIMARY KEY(room_id, room_type_id),
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (room_type_id) REFERENCES room_types(id)
);

CREATE TABLE portals ( -- Таблица порталов
    from_room_id INTEGER,
    to_room_id INTEGER,
    dimension_id INTEGER,
    target_dimension_id INTEGER,
    danger_level INTEGER,
    danger_description TEXT,
    name TEXT,
    description TEXT,
    age INTEGER,
    mean_time_to_close INTEGER,
    creator_id INTEGER,
    created_at DATETIME,
    updated_at DATETIME,
    closed BOOLEAN,
    PRIMARY KEY (from_room_id, to_room_id, dimension_id),
    FOREIGN KEY (from_room_id) REFERENCES rooms(id),
    FOREIGN KEY (to_room_id) REFERENCES rooms(id),
    FOREIGN KEY (dimension_id) REFERENCES dimensions(id),
    FOREIGN KEY (target_dimension_id) REFERENCES dimensions(id)
);

CREATE TABLE stairs ( -- Таблица лестниц
    id INTEGER PRIMARY KEY,
    name TEXT,
    description TEXT,
    danger_description TEXT,
    danger_level INTEGER,
    from_floor_id INTEGER,
    to_floor_id INTEGER,
    age INTEGER,
    created_at DATE,
    FOREIGN KEY (from_floor_id) REFERENCES floors(id),
    FOREIGN KEY (to_floor_id) REFERENCES floors(id),
    CHECK (from_floor_id <> to_floor_id)
);