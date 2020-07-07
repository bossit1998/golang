DO $$
BEGIN
    ALTER TABLE couriers
    ADD COLUMN fcm_token TEXT;
EXCEPTION
    WHEN duplicate_column THEN
        RAISE NOTICE 'Field already exists. Ignoring...';
END$$;