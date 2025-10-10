import api from "./api";

export interface ItemVenda {
  id?: number;
  id_venda: number;
  id_produto: number;
  quantidade: number;
  preco_unitario: number;
}

export async function listarItensPorVenda(idVenda: number): Promise<ItemVenda[]> {
  const res = await api.get(`/itens_venda/${idVenda}`);
  return res.data;
}

export async function adicionarItem(item: ItemVenda): Promise<void> {
  await api.post("/itens_venda", item);
}

export async function atualizarItem(id: number, item: ItemVenda): Promise<void> {
  await api.put(`/itens_venda/${id}`, item);
}

export async function deletarItem(id: number): Promise<void> {
  await api.delete(`/itens_venda/${id}`);
}
