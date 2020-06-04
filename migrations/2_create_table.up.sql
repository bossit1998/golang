CREATE TABLE IF NOT EXISTS branch_couriers (
    branch_id UUID NOT NULL,
    courier_id UUID NOT NULL REFERENCES couriers(id),
    PRIMARY KEY(branch_id, courier_id)
);