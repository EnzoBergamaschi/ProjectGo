import api from "./api";

export interface Usuario {
  id: number;
  nome: string;
  email: string;
  tipo: string;
}

interface RegisterData {
  nome: string;
  email: string;
  senha: string;
  tipo?: string;
}

export interface UpdateUsuarioData {
  nome: string;
  email: string;
  tipo: string;
  senha?: string; // opcional
}

export async function registerUser(data: RegisterData): Promise<Usuario> {
  const response = await api.post<Usuario>("/usuarios", data);
  return response.data;
}

export async function listarUsuarios(): Promise<Usuario[]> {
  const token = localStorage.getItem("token");
  const res = await api.get("/usuarios", {
    headers: { Authorization: `Bearer ${token}` },
  });
  return res.data;
}

export async function deletarUsuario(id: number): Promise<void> {
  const token = localStorage.getItem("token");
  await api.delete(`/usuarios/${id}`, {
    headers: { Authorization: `Bearer ${token}` },
  });
}

// ✅ novo: atualizar usuário (sem senha ou com senha opcional)
export async function atualizarUsuario(
  id: number,
  data: UpdateUsuarioData
): Promise<void> {
  const token = localStorage.getItem("token");
  await api.put(`/usuarios/${id}`, data, {
    headers: { Authorization: `Bearer ${token}` },
  });
}
