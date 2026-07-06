-- Múltiplas unidades prisionais de destino por reeducando

CREATE TABLE IF NOT EXISTS resocialization.inmate_destination_units (
    inmate_id      INTEGER NOT NULL REFERENCES resocialization.inmates (id) ON DELETE CASCADE,
    prison_unit_id INTEGER NOT NULL REFERENCES public.prison_units (id),
    PRIMARY KEY (inmate_id, prison_unit_id)
);

INSERT INTO resocialization.inmate_destination_units (inmate_id, prison_unit_id)
SELECT id, destination_unit_id
FROM resocialization.inmates
WHERE destination_unit_id IS NOT NULL
ON CONFLICT DO NOTHING;

ALTER TABLE resocialization.inmates
    DROP CONSTRAINT IF EXISTS inmates_destination_unit_fkey;

ALTER TABLE resocialization.inmates
    DROP COLUMN IF EXISTS destination_unit_id;
