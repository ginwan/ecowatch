CREATE TABLE readings(
    id SERIAL PRIMARY KEY,
    sensor_id INT REFERENCES sensors(id),
    value REAL,
    recorded_at TIMESTAMP,
    received_at TIMESTAMP DEFAULT now()
);