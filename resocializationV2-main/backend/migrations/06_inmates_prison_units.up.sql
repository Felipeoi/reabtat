-- Reeducandos passam a referenciar unidades prisionais (não cidades).
DELETE FROM resocialization.matches;
DELETE FROM resocialization.inmates;

ALTER TABLE resocialization.inmates
    DROP CONSTRAINT IF EXISTS inmates_origin_id_fkey,
    DROP CONSTRAINT IF EXISTS inmates_destination_id_fkey;

ALTER TABLE resocialization.inmates
    RENAME COLUMN origin_id TO origin_unit_id;

ALTER TABLE resocialization.inmates
    RENAME COLUMN destination_id TO destination_unit_id;

ALTER TABLE resocialization.inmates
    ADD CONSTRAINT inmates_origin_unit_fkey
        FOREIGN KEY (origin_unit_id) REFERENCES public.prison_units (id)
            ON UPDATE CASCADE ON DELETE RESTRICT,
    ADD CONSTRAINT inmates_destination_unit_fkey
        FOREIGN KEY (destination_unit_id) REFERENCES public.prison_units (id)
            ON UPDATE CASCADE ON DELETE RESTRICT;
