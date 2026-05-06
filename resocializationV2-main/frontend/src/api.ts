import { getToken } from "./auth";

// Em desenvolvimento, usa URL relativa para aproveitar o proxy do Vite
// Em produção, usa a variável de ambiente ou fallback
const API = import.meta.env.PROD
    ? (import.meta.env.VITE_API_URL as string) || "http://localhost:8080"
    : ""; // URL relativa em dev para usar o proxy do Vite

class ApiError extends Error {
    status?: number;
    constructor(message: string, status?: number) {
        super(message);
        this.status = status;
    }
}

async function request<T>(
    path: string,
    opts: { method?: string; body?: unknown; auth?: boolean } = {}
): Promise<T> {
    const { method = "GET", body, auth = false } = opts;
    const headers: Record<string, string> = { "Content-Type": "application/json" };
    if (auth) {
        const t = getToken();
        if (t) headers.Authorization = `Bearer ${t}`;
    }

    try {
        const res = await fetch(`${API}${path}`, {
            method,
            headers,
            body: body ? JSON.stringify(body) : undefined,
        });

        const isJson = (res.headers.get("content-type") || "").includes("application/json");
        const data = isJson ? await res.json() : await res.text();

        if (!res.ok) {
            let msg = "Erro desconhecido";
            if (isJson) {
                if (typeof data === "string") {
                    msg = data;
                } else if (data && typeof data === "object") {
                    const o = data as Record<string, unknown>;
                    msg = (typeof o.error === "string" && o.error) ||
                        (typeof o.message === "string" && o.message) ||
                        msg;
                }
            } else {
                msg = String(data);
            }
            console.error(`API Error [${res.status}] ${method} ${path}:`, msg);
            throw new ApiError(msg, res.status);
        }

        console.log(`API Success ${method} ${path}:`, data);
        return data as T;
    } catch (error) {
        if (error instanceof ApiError) {
            throw error;
        }
        console.error(`Network Error ${method} ${path}:`, error);
        throw new ApiError("Erro de conexão com o servidor. Verifique se o backend está rodando.", 0);
    }
}

/* AUTH */
export const apiSignup = (name: string, email: string, password: string, telefone?: string) =>
    request<import("./types").SignupResp>("/api/auth/signup", {
        method: "POST",
        body: { name, email, password, telefone },
    });

export const apiLogin = (email: string, password: string) =>
    request<import("./types").LoginResp>("/api/auth/login", {
        method: "POST",
        body: { email, password },
    });

export const apiMe = () =>
    request<import("./types").User>("/api/auth/me", {
        method: "GET",
        auth: true,
    });

/* USERS */
export const apiListUsers = (params: {
    limit?: number;
    offset?: number;
} = {}) => {
    const usp = new URLSearchParams();
    if (params.limit) usp.set("limit", String(params.limit));
    if (params.offset) usp.set("offset", String(params.offset));
    const q = usp.toString() ? `?${usp.toString()}` : "";
    return request<import("./types").User[]>(`/api/users${q}`, { auth: true });
};

export const apiGetUser = (id: number) =>
    request<import("./types").User>(`/api/users/${id}`, { auth: true });

export const apiCreateUser = (name: string, email: string, password: string) =>
    request<{ id: number }>(`/api/users`, {
        method: "POST",
        auth: true,
        body: { name, email, password },
    });

export const apiUpdateUser = (id: number, name: string, email: string, status: string, password?: string) =>
    request<void>(`/api/users/${id}`, {
        method: "PUT",
        auth: true,
        body: { name, email, status, password },
    });

export const apiDeleteUser = (id: number) =>
    request<void>(`/api/users/${id}`, { method: "DELETE", auth: true });

/* INMATES */
export const apiListInmates = (params: {
    limit?: number;
    offset?: number;
} = {}) => {
    const usp = new URLSearchParams();
    if (params.limit) usp.set("limit", String(params.limit));
    if (params.offset) usp.set("offset", String(params.offset));
    const q = usp.toString() ? `?${usp.toString()}` : "";
    return request<import("./types").InmatesList[]>(`/api/inmates${q}`, { auth: true });
};

export const apiGetInmate = (id: number) =>
    request<import("./types").Inmate>(`/api/inmates/${id}`, { auth: true });

export const apiCreateInmate = (data: Omit<import("./types").Inmate, "id">) =>
    request<import("./types").Inmate>(`/api/inmates`, {
        method: "POST",
        auth: true,
        body: data,
    });

export const apiUpdateInmate = (id: number, data: Omit<import("./types").Inmate, "id">) =>
    request<import("./types").Inmate>(`/api/inmates/${id}`, {
        method: "PUT",
        auth: true,
        body: data,
    });

export const apiDeleteInmate = (id: number) =>
    request<void>(`/api/inmates/${id}`, { method: "DELETE", auth: true });

/* UFS */
export const apiListUFs = () =>
    request<import("./types").UF[]>(`/api/ufs`, { auth: false });

/* CITIES */
export const apiListCities = (params: { uf_code?: string } = {}) => {
    const usp = new URLSearchParams();
    if (params.uf_code) usp.set("uf_code", params.uf_code);
    const q = usp.toString() ? `?${usp.toString()}` : "";
    return request<import("./types").City[]>(`/api/cities${q}`, { auth: false });
};

/* MATCHES */
export const apiFindMatches = () =>
    request<import("./types").MatchResult[]>(`/api/matches`, { auth: true });

export const apiFindMatchById = (myInmateId: number, matchedInmateId: number) =>
    request<import("./types").MatchResult>(`/api/matches/${myInmateId}/${matchedInmateId}`, { auth: true });
