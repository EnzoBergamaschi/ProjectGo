import { useEffect, useState } from "react";
import {
  listarProdutos,
  type Produto,
} from "../services/produtoService";
import {
  listarItensPorVenda,
  adicionarItem,
  deletarItem,
  type ItemVenda,
} from "../services/itemVendaService";

interface Props {
  idVenda: number;
  onClose: () => void;
  onAtualizarVendas: () => void; 
}

export default function ItensVendaForm({ idVenda, onClose, onAtualizarVendas }: Props) {
  const [produtos, setProdutos] = useState<Produto[]>([]);
  const [itens, setItens] = useState<ItemVenda[]>([]);
  const [novoItem, setNovoItem] = useState<ItemVenda>({
    id_venda: idVenda,
    id_produto: 0,
    quantidade: 1,
    preco_unitario: 0,
  });
  const [erro, setErro] = useState("");
  async function carregarDados() {
    try {
      const produtosData = await listarProdutos();
      setProdutos(Array.isArray(produtosData) ? produtosData : []);

      const itensData = await listarItensPorVenda(idVenda);
      setItens(Array.isArray(itensData) ? itensData : []);
    } catch (err) {
      console.error("Erro ao carregar dados:", err);
      setErro("Erro ao carregar dados dos itens.");
      setItens([]);
    }
  }

  useEffect(() => {
    carregarDados();
  }, [idVenda]);
  async function handleAdd() {
    setErro("");

    if (!novoItem.id_produto || novoItem.quantidade <= 0) {
      setErro("Selecione um produto e insira uma quantidade válida.");
      return;
    }

    try {
      const produto = produtos.find((p) => p.id === novoItem.id_produto);
      if (!produto) {
        setErro("Produto inválido.");
        return;
      }

      await adicionarItem({
        ...novoItem,
        preco_unitario: produto.preco,
      });

      setNovoItem({
        id_venda: idVenda,
        id_produto: 0,
        quantidade: 1,
        preco_unitario: 0,
      });

      await carregarDados();
      await onAtualizarVendas(); 
    } catch (err) {
      console.error(err);
      setErro("Erro ao adicionar item.");
    }
  }
  async function handleDelete(id: number) {
    try {
      await deletarItem(id);
      await carregarDados();
      await onAtualizarVendas();
    } catch (err) {
      console.error(err);
      setErro("Erro ao excluir item.");
    }
  }
  function calcularTotal(): number {
    if (!Array.isArray(itens) || itens.length === 0) return 0;
    return itens.reduce((acc, item) => {
      const qtd = Number(item.quantidade) || 0;
      const preco = Number(item.preco_unitario) || 0;
      return acc + qtd * preco;
    }, 0);
  }

  const total = calcularTotal();

  function handleFechar() {
    setItens([]);
    setNovoItem({
      id_venda: idVenda,
      id_produto: 0,
      quantidade: 1,
      preco_unitario: 0,
    });
    onClose();
  }
  return (
    <div className="bg-slate-800 p-6 rounded-xl shadow-lg mt-6">
      <div className="flex justify-between items-center mb-4">
        <h2 className="text-xl font-semibold">Itens da Venda #{idVenda}</h2>
        <div className="flex gap-2">
          <button
            onClick={handleFechar}
            className="bg-slate-600 hover:bg-slate-500 px-3 py-1 rounded transition"
          >
            Fechar
          </button>
        </div>
      </div>

      {erro && <p className="text-red-400 mb-2">{erro}</p>}

      <div className="flex gap-2 mb-4">
        <select
          value={novoItem.id_produto}
          onChange={(e) =>
            setNovoItem({ ...novoItem, id_produto: Number(e.target.value) })
          }
          className="flex-1 p-2 rounded bg-slate-700 focus:outline-none"
        >
          <option value={0}>Selecione um produto</option>
          {produtos.map((p) => (
            <option key={p.id} value={p.id}>
              {p.nome} - R$ {p.preco.toFixed(2)}
            </option>
          ))}
        </select>

        <input
          type="number"
          placeholder="Qtd"
          value={novoItem.quantidade}
          min={1}
          onChange={(e) =>
            setNovoItem({ ...novoItem, quantidade: Number(e.target.value) })
          }
          className="w-24 p-2 rounded bg-slate-700 text-center"
        />

        <button
          onClick={handleAdd}
          className="bg-blue-600 hover:bg-blue-500 px-4 py-2 rounded font-semibold transition"
        >
          Adicionar
        </button>
      </div>

      <table className="w-full border-collapse bg-slate-900 rounded-lg shadow">
        <thead>
          <tr className="bg-slate-700">
            <th className="p-2 text-left">Produto</th>
            <th className="p-2 text-left">Qtd</th>
            <th className="p-2 text-left">Preço Unit.</th>
            <th className="p-2 text-left">Subtotal</th>
            <th className="p-2 text-center">Ações</th>
          </tr>
        </thead>
        <tbody>
          {Array.isArray(itens) && itens.length > 0 ? (
            itens.map((i) => {
              const produto = produtos.find((p) => p.id === i.id_produto);
              return (
                <tr key={i.id} className="border-t border-slate-700">
                  <td className="p-2">{produto?.nome || "Desconhecido"}</td>
                  <td className="p-2">{i.quantidade}</td>
                  <td className="p-2">R$ {i.preco_unitario.toFixed(2)}</td>
                  <td className="p-2">
                    R$ {(i.quantidade * i.preco_unitario).toFixed(2)}
                  </td>
                  <td className="p-2 text-center">
                    <button
                      onClick={() => handleDelete(i.id!)}
                      className="bg-red-600 hover:bg-red-500 px-3 py-1 rounded font-semibold transition"
                    >
                      Excluir
                    </button>
                  </td>
                </tr>
              );
            })
          ) : (
            <tr>
              <td colSpan={5} className="text-center p-4 text-gray-400 italic">
                Nenhum item adicionado.
              </td>
            </tr>
          )}
        </tbody>
      </table>

      <div className="text-right mt-4 font-semibold text-lg">
        Total: R$ {total.toFixed(2)}
      </div>
    </div>
  );
}
