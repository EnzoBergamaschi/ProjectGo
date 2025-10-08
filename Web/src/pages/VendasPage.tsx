import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Navbar from "../components/layout/Navbar";
import {
  listarVendas,
  criarVenda,
  atualizarVenda,
  deletarVenda,
  type Venda,
} from "../services/vendaService";
import {
  listarUsuarios,
  type RegisterResponse as Usuario,
} from "../services/userService";
import ItensVendaForm from "./ItensVendaForm";

export default function VendasPage() {
  const [vendas, setVendas] = useState<Venda[]>([]);
  const [usuarios, setUsuarios] = useState<Usuario[]>([]);
  const [novaVenda, setNovaVenda] = useState<Venda>({
    id_usuario: 0,
    status: "pendente",
    total: 0,
  });
  const [editando, setEditando] = useState<Venda | null>(null);
  const [erro, setErro] = useState("");
  const navigate = useNavigate();

  // ============================================================
  // Carregar vendas e usuários
  // ============================================================
  async function carregarVendas() {
    try {
      const vendasData = await listarVendas();
      setVendas(vendasData);
    } catch (err) {
      console.error("Erro ao carregar vendas:", err);
      setErro("Erro ao carregar vendas.");
    }
  }

  async function carregarUsuarios() {
    try {
      const usuariosData = await listarUsuarios();
      setUsuarios(usuariosData);
    } catch (err) {
      console.error("Erro ao carregar usuários:", err);
      setErro("Erro ao carregar usuários.");
    }
  }

  useEffect(() => {
    Promise.all([carregarVendas(), carregarUsuarios()]);
  }, []);

  // ============================================================
  // Enviar formulário (criar ou atualizar venda)
  // ============================================================
  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setErro("");

    if (!novaVenda.id_usuario) {
      setErro("Selecione um usuário antes de registrar a venda.");
      return;
    }

    try {
      if (editando) {
        await atualizarVenda(editando.id!, novaVenda);
      } else {
        await criarVenda(novaVenda);
      }

      setNovaVenda({ id_usuario: 0, status: "pendente", total: 0 });
      setEditando(null);
      await carregarVendas();
    } catch (err) {
      console.error("Erro ao salvar venda:", err);
      setErro("Erro ao salvar venda.");
    }
  }

  // ============================================================
  // Excluir venda
  // ============================================================
  async function handleDelete(id: number) {
    if (!confirm("Deseja realmente excluir esta venda?")) return;

    try {
      await deletarVenda(id);
      await carregarVendas();
    } catch (err) {
      console.error("Erro ao excluir venda:", err);
      setErro("Erro ao excluir venda.");
    }
  }

  // ============================================================
  // Render
  // ============================================================
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

        <h1 className="text-3xl font-bold mb-4">Gerenciar Vendas</h1>
        {erro && <p className="text-red-400 mb-4">{erro}</p>}

        {/* Formulário de criação/edição */}
        <form
          onSubmit={handleSubmit}
          className="bg-slate-800 p-6 rounded-lg mb-8 shadow-lg space-y-3"
        >
          <h2 className="text-xl font-semibold mb-2">
            {editando ? "Editar Venda" : "Nova Venda"}
          </h2>

          {/* Selecionar usuário */}
          <select
            value={novaVenda.id_usuario}
            onChange={(e) =>
              setNovaVenda({
                ...novaVenda,
                id_usuario: Number(e.target.value),
              })
            }
            className="w-full p-2 rounded bg-slate-700 focus:outline-none"
            required
          >
            <option value={0}>Selecione o usuário</option>
            {usuarios.map((u) => (
              <option key={u.id} value={u.id}>
                {u.nome} ({u.email})
              </option>
            ))}
          </select>

          {/* Status */}
          <select
            value={novaVenda.status}
            onChange={(e) =>
              setNovaVenda({ ...novaVenda, status: e.target.value })
            }
            className="w-full p-2 rounded bg-slate-700 focus:outline-none"
          >
            <option value="pendente">Pendente</option>
            <option value="pago">Pago</option>
            <option value="enviado">Enviado</option>
            <option value="cancelado">Cancelado</option>
          </select>

          <button
            type="submit"
            className="bg-blue-600 hover:bg-blue-500 px-4 py-2 rounded font-semibold transition"
          >
            {editando ? "Salvar Alterações" : "Registrar Venda"}
          </button>
        </form>

        {/* Tabela de vendas */}
        <table className="w-full border-collapse bg-slate-800 rounded-lg shadow-lg">
          <thead>
            <tr className="bg-slate-700">
              <th className="p-3 text-left">ID</th>
              <th className="p-3 text-left">Usuário</th>
              <th className="p-3 text-left">Status</th>
              <th className="p-3 text-left">Total</th>
              <th className="p-3 text-center">Ações</th>
            </tr>
          </thead>
          <tbody>
            {vendas.map((v) => (
              <tr key={v.id} className="border-t border-slate-600">
                <td className="p-3">{v.id}</td>
                <td className="p-3">{v.usuario_nome || "—"}</td>
                <td className="p-3 capitalize">{v.status}</td>
                <td className="p-3">R$ {v.total.toFixed(2)}</td>
                <td className="p-3 flex gap-2 justify-center">
                  <button
                    onClick={() => {
                      setEditando(v);
                      setNovaVenda({
                        id_usuario: v.id_usuario,
                        status: v.status,
                        total: v.total,
                      });
                    }}
                    className="bg-yellow-500 hover:bg-yellow-400 px-3 py-1 rounded text-black font-semibold"
                  >
                    Editar
                  </button>

                  <button
                    onClick={() => handleDelete(v.id!)}
                    className="bg-red-600 hover:bg-red-500 px-3 py-1 rounded font-semibold"
                  >
                    Excluir
                  </button>
                </td>
              </tr>
            ))}

            {vendas.length === 0 && (
              <tr>
                <td colSpan={5} className="text-center p-4 text-gray-400">
                  Nenhuma venda registrada.
                </td>
              </tr>
            )}
          </tbody>
        </table>

        {/* Modal de Itens da Venda */}
        {editando && (
          <ItensVendaForm
            idVenda={editando.id!}
            onClose={() => setEditando(null)}
            onAtualizarVendas={carregarVendas} // 🔹 atualização automática de totais
          />
        )}
      </main>
    </div>
  );
}
