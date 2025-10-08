import api from "./api";

interface LoginData {
  email: string;
  senha: string;
}

interface LoginResponse {
  token: string;
}

// Faz login e retorna o token JWT
export async function loginUser(data: LoginData): Promise<LoginResponse> {
  const response = await api.post<LoginResponse>("/login", data);
  return response.data;
}

// Exemplo de função auxiliar para logout
export function logoutUser() {
  localStorage.removeItem("token");
}

// Exemplo de verificação de login
export function isAuthenticated(): boolean {
  return !!localStorage.getItem("token");
}
