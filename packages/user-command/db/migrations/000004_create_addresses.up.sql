CREATE TABLE countries (
    id INT PRIMARY KEY,
    name TEXT NOT NULL,
    code VARCHAR(10) UNIQUE
);

CREATE TABLE cities (
    id INT PRIMARY KEY,
    country_id INT REFERENCES countries(id),
    name TEXT NOT NULL,
    type TEXT
);

CREATE TABLE districts (
    id INT PRIMARY KEY,
    city_id INT REFERENCES cities(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    type TEXT
);

CREATE TABLE wards (
    id INT PRIMARY KEY,
    district_id INT REFERENCES districts(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    type TEXT
);

CREATE INDEX idx_cities_country_id ON cities(country_id);
CREATE INDEX idx_districts_city_id ON districts(city_id);
CREATE INDEX idx_wards_district_id ON wards(district_id);