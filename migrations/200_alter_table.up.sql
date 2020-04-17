
DROP TABLE IF EXISTS courier_details;
DROP TABLE IF EXISTS distributor_fare;
DROP TABLE IF EXISTS courier_fare;
DROP TABLE IF EXISTS courier_vehicle;
DROP TABLE IF EXISTS courier;
DROP TABLE IF EXISTS distributor;


CREATE TABLE IF NOT EXISTS distributor (
    id uuid PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS courier (
    id uuid PRIMARY KEY,
    distributor_id uuid REFERENCES distributor(id) ON DELETE CASCADE,
    phone VARCHAR(100) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS courier_details (
    courier_id uuid PRIMARY KEY REFERENCES courier(id) ON DELETE CASCADE,
    passport_number VARCHAR(100) NOT NULL,
    gender VARCHAR(100),
    birth_date TIMESTAMP NOT NULL,
    address VARCHAR(100),
    img VARCHAR(100),
    lisense_number VARCHAR(100) NOT NULL,
    lisense_given_date TIMESTAMP NOT NULL,
    lisense_expiry_date TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS courier_vehicle (
    id uuid PRIMARY KEY,
    courier_id uuid REFERENCES courier(id) ON DELETE CASCADE,
    model VARCHAR(100) NOT NULL,
    vehicle_number VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status BOOLEAN NOT NULL DEFAULT TRUE
);
