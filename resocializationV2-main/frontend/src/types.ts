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

export type InmatesResponsible = {
    attorney: string;
    phone: string;
};

export type Inmate = {
    id: number;
    origin_id: number;
    origin?: City;
    custody: "CLOSED" | "SEMI_OPEN" | "OPEN";
    destination_id: number;
    destination?: City;
    responsible: InmatesResponsible;
};

export type InmatesList = {
    id: number;
    origin_id: number;
    origin?: City;
    destination_id: number;
    destination?: City;
    custody: string;
};

export type LoginResp = { token: string };
export type SignupResp = { user_id: number };

// Match types
export type InmateMatchInfo = {
    id: number;
    origin_id: number;
    origin?: City;
    destination_id: number;
    destination?: City;
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
