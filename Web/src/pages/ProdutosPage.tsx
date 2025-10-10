import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Navbar from "../components/layout/Navbar";
import {
  listarProdutos,
  criarProduto,
  atualizarProduto,
  deletarProduto,
  type Produto,
} from "../services/produtoService";

export default function ProdutosPage() {
  const [produtos, setProdutos] = useState<Produto[]>([]);
  const [novoProduto, setNovoProduto] = useState({
    nome: "",
    preco: "",
    estoque: "",
  });
  const [editando, setEditando] = useState<Produto | null>(null);
  const [erro, setErro] = useState("");
  const navigate = useNavigate();

  async function carregarProdutos() {
    try {
      const data = await listarProdutos();
      setProdutos(data);
    } catch (err: any) {
      console.error(err);
      setErro("Erro ao carregar produtos (verifique o token ou permissões).");
    }
  }

  useEffect(() => {
    carregarProdutos();
  }, []);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    try {
      const produtoFormatado = {
        nome: novoProduto.nome,
        preco: parseFloat(novoProduto.preco) || 0,
        estoque: parseInt(novoProduto.estoque) || 0,
      };

      if (editando) {
        await atualizarProduto(editando.id!, produtoFormatado);
      } else {
        await criarProduto(produtoFormatado);
      }

      setNovoProduto({ nome: "", preco: "", estoque: "" });
      setEditando(null);
      carregarProdutos();
    } catch (err) {
      console.error(err);
      setErro("Erro ao salvar produto.");
    }
  }

  async function handleDelete(id: number) {
    if (!confirm("Deseja realmente excluir este produto?")) return;
    try {
      await deletarProduto(id);
      carregarProdutos();
    } catch (err) {
      console.error(err);
      setErro("Erro ao excluir produto.");
    }
  }

  return (
    <div className="flex flex-col min-h-screen bg-slate-900 text-white">
      <Navbar onLogout={() => navigate("/")} />

      <main className="flex-1 p-8">
        <button
          onClick={() => navigate("/dashboard")}
          className="mb-6 bg-slate-700 hover:bg-slate-600 px-4 py-2 rounded-lg font-semibold transition text-sm"
        >
          ← Voltar ao Dashboard
        </button>

        <h1 className="text-3xl font-bold mb-4">Gerenciar Produtos</h1>
        {erro && <p className="text-red-400 mb-4">{erro}</p>}

        <form
          onSubmit={handleSubmit}
          className="bg-slate-800 p-6 rounded-lg mb-8 shadow-lg space-y-4"
        >
          <h2 className="text-xl font-semibold mb-2">
            {editando ? "Editar Produto" : "Novo Produto"}
          </h2>

          <div>
            <label className="block mb-1 text-sm font-semibold text-gray-300">
              Nome do Produto
            </label>
            <input
              type="text"
              placeholder="Ex: Teclado Mecânico RGB"
              value={novoProduto.nome}
              onChange={(e) =>
                setNovoProduto({ ...novoProduto, nome: e.target.value })
              }
              className="w-full p-2 rounded bg-slate-700 focus:outline-none"
              required
            />
          </div>

          <div>
            <label className="block mb-1 text-sm font-semibold text-gray-300">
              Preço (em R$)
            </label>
            <input
              type="number"
              step="0.01"
              placeholder="Ex: 199.90"
              value={novoProduto.preco}
              onChange={(e) =>
                setNovoProduto({ ...novoProduto, preco: e.target.value })
              }
              className="w-full p-2 rounded bg-slate-700 focus:outline-none"
              required
            />
          </div>

          <div>
            <label className="block mb-1 text-sm font-semibold text-gray-300">
              Quantidade em Estoque
            </label>
            <input
              type="number"
              placeholder="Ex: 10"
              value={novoProduto.estoque}
              onChange={(e) =>
                setNovoProduto({ ...novoProduto, estoque: e.target.value })
              }
              className="w-full p-2 rounded bg-slate-700 focus:outline-none"
              required
            />
          </div>

          <button
            type="submit"
            className={`${
              editando
                ? "bg-yellow-500 hover:bg-yellow-400 text-black"
                : "bg-blue-600 hover:bg-blue-500 text-white"
            } px-4 py-2 rounded font-semibold transition`}
          >
            {editando ? "Salvar Alterações" : "Adicionar Produto"}
          </button>
        </form>

        <table className="w-full border-collapse bg-slate-800 rounded-lg shadow-lg">
          <thead>
            <tr className="bg-slate-700">
              <th className="p-3 text-left">ID</th>
              <th className="p-3 text-left">Nome</th>
              <th className="p-3 text-left">Preço</th>
              <th className="p-3 text-left">Estoque</th>
              <th className="p-3 text-center">Ações</th>
            </tr>
          </thead>
          <tbody>
            {produtos.map((p) => (
              <tr key={p.id} className="border-t border-slate-600">
                <td className="p-3">{p.id}</td>
                <td className="p-3">{p.nome}</td>
                <td className="p-3">R$ {p.preco.toFixed(2)}</td>
                <td className="p-3">{p.estoque}</td>
                <td className="p-3 flex gap-2 justify-center">
                  <button
                    onClick={() => {
                      setEditando(p);
                      setNovoProduto({
                        nome: p.nome,
                        preco: p.preco.toString(),
                        estoque: p.estoque.toString(),
                      });
                    }}
                    className="bg-yellow-500 hover:bg-yellow-400 px-3 py-1 rounded text-black font-semibold"
                  >
                    Editar
                  </button>

                  <button
                    onClick={() => handleDelete(p.id!)}
                    className="bg-red-600 hover:bg-red-500 px-3 py-1 rounded font-semibold"
                  >
                    Excluir
                  </button>
                </td>
              </tr>
            ))}

            {produtos.length === 0 && (
              <tr>
                <td colSpan={5} className="text-center p-4 text-gray-400">
                  Nenhum produto encontrado.
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </main>
    </div>
  );
}
