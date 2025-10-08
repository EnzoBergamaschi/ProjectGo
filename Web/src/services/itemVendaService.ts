import api from "./api";

export interface ItemVenda {
  id?: number;
  id_venda: number;
  id_produto: number;
  quantidade: number;
  preco_unitario: number;
}

// Listar todos os itens de uma venda específica
export async function listarItensPorVenda(idVenda: number): Promise<ItemVenda[]> {
  const res = await api.get(`/itens_venda/${idVenda}`);
  return res.data;
}

// Criar um novo item de venda
export async function adicionarItem(item: ItemVenda): Promise<void> {
  await api.post("/itens_venda", item);
}

// Atualizar um item de venda
export async function atualizarItem(id: number, item: ItemVenda): Promise<void> {
  await api.put(`/itens_venda/${id}`, item);
}

// Deletar item
export async function deletarItem(id: number): Promise<void> {
  await api.delete(`/itens_venda/${id}`);
}
