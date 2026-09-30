CREATE TABLE alerts(
    id SERIAL PRIMARY KEY,
    sensor_id INT REFERENCES sensors(id),
    status TEXT,
    value REAL,
    message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- docker compose exec -T db psql -U postgres -d ecowatch < db/migrations/005_create_alerts.sql
-- docker compose exec db psql -U postgres -d ecowatch -c "\dt"