CREATE TABLE IF NOT EXISTS courier_vendors (
    vendor_id UUID NOT NULL,
    courier_id UUID NOT NULL REFERENCES couriers(id),
    PRIMARY KEY(vendor_id, courier_id)
);