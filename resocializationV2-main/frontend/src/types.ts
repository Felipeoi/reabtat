export type User = {
    id: number;
    name: string;
    email: string;
    status: string;
    role: string;
    telefone?: string;
    created_at?: string;
    updated_at?: string;
};

export type UF = {
    id: number;
    code: string;
    name: string;
};

export type City = {
    id: number;
    ibge_code: string;
    name: string;
    uf_code: string;
};

export type PrisonUnit = {
    id: number;
    name: string;
    uf_code: string;
};

export type InmatesResponsible = {
    attorney: string;
    phone: string;
};

export type Inmate = {
    id: number;
    origin_unit_id: number;
    origin_unit?: PrisonUnit;
    custody: "CLOSED" | "SEMI_OPEN" | "OPEN";
    destination_unit_ids: number[];
    destination_units?: PrisonUnit[];
    responsible: InmatesResponsible;
};

export type InmatesList = {
    id: number;
    origin_unit_id: number;
    origin_unit?: PrisonUnit;
    destination_unit_ids: number[];
    destination_units?: PrisonUnit[];
    custody: string;
};

export type LoginResp = { token: string };
export type SignupResp = { user_id: number };

export type InmateMatchInfo = {
    id: number;
    origin_unit_id: number;
    origin_unit?: PrisonUnit;
    destination_unit_ids: number[];
    destination_units?: PrisonUnit[];
    custody: string;
    user_id: number;
    responsible: InmatesResponsible;
};

export type MatchResult = {
    my_inmate: InmateMatchInfo;
    matched_inmate: InmateMatchInfo;
    match_score: number;
    custody: string;
};
