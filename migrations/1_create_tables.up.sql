CREATE TABLE IF NOT EXISTS distributors (
     id uuid PRIMARY KEY,
     access_token VARCHAR NOT NULL UNIQUE,
     name VARCHAR(100) NOT NULL UNIQUE,
     phone VARCHAR(100) NOT NULL UNIQUE,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP,
     deleted_at TIMESTAMP,
     is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS parks (
     id uuid PRIMARY KEY,
     distributor_id uuid NOT NULL REFERENCES distributors(id),
     name VARCHAR(100) NOT NULL UNIQUE,
     location GEOMETRY NOT NULL,
     address VARCHAR,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP,
     deleted_at TIMESTAMP,
     is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS couriers (
     id uuid PRIMARY KEY,
     access_token VARCHAR NOT NULL UNIQUE,
     distributor_id uuid REFERENCES distributors(id),
     park_id uuid REFERENCES parks(id),
     phone VARCHAR(100) NOT NULL,
     first_name VARCHAR(100) NOT NULL,
     last_name VARCHAR(100) NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP,
     deleted_at TIMESTAMP,
     is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS courier_details (
     courier_id uuid PRIMARY KEY REFERENCES couriers(id),
     passport_number VARCHAR(100) NOT NULL,
     gender VARCHAR(100),
     birth_date TIMESTAMP NOT NULL,
     address VARCHAR(100),
     img VARCHAR(100),
     lisense_number VARCHAR(100) NOT NULL,
     lisense_given_date TIMESTAMP NOT NULL,
     lisense_expiry_date TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS courier_vehicles (
     id uuid PRIMARY KEY,
     courier_id uuid NOT NULL REFERENCES couriers(id),
     model VARCHAR(100) NOT NULL,
     vehicle_number VARCHAR(100) NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP,
     deleted_at TIMESTAMP,
     is_active BOOLEAN NOT NULL DEFAULT TRUE
);
