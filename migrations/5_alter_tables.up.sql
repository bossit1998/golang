drop index couriers_u1;

do $$
    begin
        alter table couriers alter column deleted_at type integer using date_part('epoch', deleted_at)::int;
        alter table couriers alter column deleted_at set default 0;
        alter table couriers add unique (phone, shipper_id, deleted_at);
    exception
        when duplicate_column then
            RAISE NOTICE 'Already existed';
    end $$;

update couriers set deleted_at=0 where deleted_at is null;