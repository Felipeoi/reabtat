CREATE TABLE IF NOT EXISTS public.prison_units
(
    id      SERIAL PRIMARY KEY,
    name    TEXT    NOT NULL,
    uf_code CHAR(2) NOT NULL REFERENCES public.ufs (code) ON UPDATE CASCADE,
    CONSTRAINT prison_units_uf_name_uk UNIQUE (uf_code, name)
);

CREATE INDEX IF NOT EXISTS idx_prison_units_uf_code ON public.prison_units (uf_code);
