import api from "./api";

export interface Venda {
  id?: number;
  id_usuario: number;
  usuario_nome?: string;
  status: string;
  total: number;
}

export async function listarVendas(): Promise<Venda[]> {
  const res = await api.get("/vendas");
  return res.data;
}

export async function criarVenda(venda: Venda) {
  await api.post("/vendas", venda);
}

export async function atualizarVenda(id: number, venda: Venda) {
  await api.put(`/vendas/${id}`, venda);
}

export async function deletarVenda(id: number) {
  await api.delete(`/vendas/${id}`);
}
