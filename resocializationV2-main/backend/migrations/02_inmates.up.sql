-- Tabela de UFs
CREATE TABLE IF NOT EXISTS public.ufs (
    id SERIAL PRIMARY KEY,
    code CHAR(2)  NOT NULL UNIQUE, -- ex.: 'SP', 'RJ'
    name TEXT     NOT NULL         -- ex.: 'São Paulo'
);

CREATE TABLE IF NOT EXISTS public.cities
(
    id        bigserial PRIMARY KEY,
    ibge_code char(7)      NOT NULL,
    name      varchar(255) NOT NULL,
    uf_code   char(2)      NOT NULL REFERENCES public.ufs (code)
    ON UPDATE CASCADE,
    CONSTRAINT cities_ibge_code_uk UNIQUE (ibge_code),
    CONSTRAINT cities_uf_name_uk UNIQUE (uf_code, name)
    );

-- Tabela de Inmates (cadastro de presos)
CREATE TABLE IF NOT EXISTS resocialization.inmates
(
    id                   SERIAL PRIMARY KEY,
    custody              VARCHAR(16)  NOT NULL CHECK (custody IN ('CLOSED', 'SEMI_OPEN', 'OPEN')),
    origin_id               INTEGER NOT NULL REFERENCES public.cities (id) ON UPDATE CASCADE ON DELETE RESTRICT,
    destination_id INTEGER NOT NULL REFERENCES public.cities (id) ON UPDATE CASCADE ON DELETE RESTRICT,
    responsible_attorney VARCHAR(100),
    responsible_phone    VARCHAR(20),
    user_id              INTEGER      NOT NULL REFERENCES resocialization.users (id) ON UPDATE CASCADE ON DELETE RESTRICT
);