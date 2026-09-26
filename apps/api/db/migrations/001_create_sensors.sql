CREATE TABLE sensors(
    id SERIAL PRIMARY KEY,
    name TEXT,
    type TEXT,
    facility_id INTEGER,
    unit TEXT,
    min_threshold REAL,
    max_threshold REAL,
    status TEXT
);