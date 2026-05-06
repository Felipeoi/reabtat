const KEY = "auth_token";

export function setToken(token: string) {
    localStorage.setItem(KEY, token);
}
export function getToken(): string {
    return localStorage.getItem(KEY) || "";
}
export function isAuthed(): boolean {
    return !!getToken();
}
export function logout() {
    localStorage.removeItem(KEY);
}
