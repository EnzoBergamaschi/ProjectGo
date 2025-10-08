import api from "./api";

export interface Produto {
  id?: number;
  nome: string;
  preco: number;
  estoque: number;
}

export async function listarProdutos(): Promise<Produto[]> {
  const res = await api.get("/produtos");
  return res.data;
}

export async function criarProduto(produto: Produto) {
  await api.post("/produtos", produto);
}

export async function atualizarProduto(id: number, produto: Produto) {
  await api.put(`/produtos/${id}`, produto);
}

export async function deletarProduto(id: number) {
  await api.delete(`/produtos/${id}`);
}
