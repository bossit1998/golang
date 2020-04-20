CREATE TABLE IF NOT EXISTS parks (
     id uuid PRIMARY KEY,
     distributor_id uuid NOT NULL REFERENCES distributors(id),
     name VARCHAR(100) NOT NULL,
     location GEOMETRY NOT NULL,
     address VARCHAR,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP,
     deleted_at TIMESTAMP,
     is_active BOOLEAN NOT NULL DEFAULT TRUE
);

ALTER TABLE couriers
	add park_id uuid REFERENCES parks(id);

