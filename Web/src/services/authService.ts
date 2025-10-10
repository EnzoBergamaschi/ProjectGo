import api from "./api";

interface LoginData {
  email: string;
  senha: string;
}

interface LoginResponse {
  token: string;
}

export async function loginUser(data: LoginData): Promise<LoginResponse> {
  const response = await api.post<LoginResponse>("/login", data);
  return response.data;
}

export function logoutUser() {
  localStorage.removeItem("token");
}

export function isAuthenticated(): boolean {
  return !!localStorage.getItem("token");
}

export function getUserRole(): string | null {
  const token = localStorage.getItem("token");
  if (!token) return null;

  try {
    const payload = JSON.parse(atob(token.split(".")[1]));
    return payload.tipo || payload.role || null;
  } catch {
    return null;
  }
}
