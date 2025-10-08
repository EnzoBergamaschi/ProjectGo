import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Navbar from "../components/layout/Navbar";
import {
  listarUsuarios,
  deletarUsuario,
  registerUser,
  atualizarUsuario,
  type Usuario,
} from "../services/userService";

export default function UsuariosPage() {
  const [usuarios, setUsuarios] = useState<Usuario[]>([]);
  const [loading, setLoading] = useState(true);
  const [erro, setErro] = useState("");

  // feedback visual
  const [mensagem, setMensagem] = useState<string | null>(null);

  // modal criar
  const [showCreate, setShowCreate] = useState(false);
  const [novoUsuario, setNovoUsuario] = useState({
    nome: "",
    email: "",
    senha: "",
    tipo: "cliente",
  });

  // modal editar
  const [showEdit, setShowEdit] = useState(false);
  const [editUserId, setEditUserId] = useState<number | null>(null);
  const [formEdicao, setFormEdicao] = useState({
    nome: "",
    email: "",
    tipo: "cliente",
    senha: "", // opcional (se quiser trocar)
  });

  const navigate = useNavigate();

  useEffect(() => {
    carregarUsuarios();
  }, []);

  async function carregarUsuarios() {
    try {
      const data = await listarUsuarios();
      setUsuarios(data);
    } catch (error) {
      console.error("Erro ao listar usuários:", error);
      setErro("Erro ao carregar usuários.");
    } finally {
      setLoading(false);
    }
  }

  async function handleDelete(id: number) {
    if (!confirm("Tem certeza que deseja excluir este usuário?")) return;
    try {
      await deletarUsuario(id);
      setUsuarios((prev) => prev.filter((u) => u.id !== id));
      setMensagem("Usuário excluído com sucesso!");
      setTimeout(() => setMensagem(null), 3000);
    } catch (error) {
      alert("Erro ao excluir usuário");
      console.error(error);
    }
  }

  async function handleCreate(e: React.FormEvent) {
    e.preventDefault();
    try {
      await registerUser(novoUsuario);
      setShowCreate(false);
      setNovoUsuario({ nome: "", email: "", senha: "", tipo: "cliente" });
      await carregarUsuarios();

      setMensagem("Usuário criado com sucesso!");
      setTimeout(() => setMensagem(null), 3000);
    } catch (error) {
      console.error(error);
      alert("Erro ao criar usuário");
    }
  }

  function abrirEdicao(u: Usuario) {
    setEditUserId(u.id);
    setFormEdicao({ nome: u.nome, email: u.email, tipo: u.tipo, senha: "" });
    setShowEdit(true);
  }

  async function handleEdit(e: React.FormEvent) {
    e.preventDefault();
    if (editUserId == null) return;

    try {
      const payload =
        formEdicao.senha.trim() === ""
          ? { nome: formEdicao.nome, email: formEdicao.email, tipo: formEdicao.tipo }
          : {
              nome: formEdicao.nome,
              email: formEdicao.email,
              tipo: formEdicao.tipo,
              senha: formEdicao.senha,
            };

      await atualizarUsuario(editUserId, payload);
      setShowEdit(false);
      await carregarUsuarios();

      // ✅ Mensagem de sucesso
      setMensagem("Usuário atualizado com sucesso!");
      setTimeout(() => setMensagem(null), 3000);
    } catch (error) {
      console.error(error);
      alert("Erro ao atualizar usuário");
    }
  }

  if (loading)
    return (
      <div className="flex items-center justify-center h-screen text-gray-300">
        Carregando usuários...
      </div>
    );

  return (
    <div className="flex flex-col min-h-screen bg-slate-900 text-white">
      <Navbar onLogout={() => navigate("/")} />

      {/* ✅ Toast simples */}
      {mensagem && (
        <div className="fixed top-4 right-4 bg-green-600 text-white px-4 py-2 rounded-lg shadow-lg animate-fade-in">
          {mensagem}
        </div>
      )}

      <main className="flex-1 p-8">
        {/* topo padronizado */}
        <div className="flex justify-between items-center mb-6">
          <button
            onClick={() => navigate("/dashboard")}
            className="bg-slate-700 hover:bg-slate-600 px-4 py-2 rounded-lg font-semibold transition text-sm"
          >
            ← Voltar ao Dashboard
          </button>

          <button
            onClick={() => setShowCreate(true)}
            className="bg-green-600 hover:bg-green-500 px-4 py-2 rounded font-semibold transition text-sm"
          >
            + Adicionar Usuário
          </button>
        </div>

        <h1 className="text-3xl font-bold mb-6">Usuários cadastrados</h1>
        {erro && <p className="text-red-400 mb-4">{erro}</p>}

        {/* tabela */}
        <div className="bg-slate-800 rounded-lg shadow-lg overflow-hidden">
          <table className="w-full text-sm text-left border-collapse">
            <thead className="bg-slate-700 text-gray-300">
              <tr>
                <th className="px-6 py-3 font-semibold">ID</th>
                <th className="px-6 py-3 font-semibold">Nome</th>
                <th className="px-6 py-3 font-semibold">Email</th>
                <th className="px-6 py-3 font-semibold">Tipo</th>
                <th className="px-6 py-3 font-semibold text-center">Ações</th>
              </tr>
            </thead>
            <tbody>
              {usuarios.map((u) => (
                <tr
                  key={u.id}
                  className="border-t border-slate-700 hover:bg-slate-700/50 transition"
                >
                  <td className="px-6 py-3">{u.id}</td>
                  <td className="px-6 py-3">{u.nome}</td>
                  <td className="px-6 py-3">{u.email}</td>
                  <td className="px-6 py-3 capitalize">{u.tipo}</td>
                  <td className="px-6 py-3">
                    <div className="flex gap-2 justify-center">
                      <button
                        onClick={() => abrirEdicao(u)}
                        className="bg-yellow-500 hover:bg-yellow-400 px-3 py-1 rounded text-black font-semibold"
                      >
                        Editar
                      </button>
                      <button
                        onClick={() => handleDelete(u.id)}
                        className="bg-red-600 hover:bg-red-500 text-white px-3 py-1 rounded font-semibold"
                      >
                        Excluir
                      </button>
                    </div>
                  </td>
                </tr>
              ))}

              {usuarios.length === 0 && (
                <tr>
                  <td colSpan={5} className="text-center p-4 text-gray-400">
                    Nenhum usuário encontrado.
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
      </main>

      {/* Modal Criar */}
      {showCreate && (
        <div className="fixed inset-0 bg-black/60 flex items-center justify-center z-50">
          <div className="bg-slate-800 p-6 rounded-lg shadow-lg w-[420px]">
            <h2 className="text-xl font-semibold mb-4">Novo Usuário</h2>
            <form onSubmit={handleCreate} className="space-y-3">
              <input
                type="text"
                placeholder="Nome"
                className="w-full p-2 rounded bg-slate-700 focus:outline-none"
                value={novoUsuario.nome}
                onChange={(e) =>
                  setNovoUsuario({ ...novoUsuario, nome: e.target.value })
                }
                required
              />
              <input
                type="email"
                placeholder="Email"
                className="w-full p-2 rounded bg-slate-700 focus:outline-none"
                value={novoUsuario.email}
                onChange={(e) =>
                  setNovoUsuario({ ...novoUsuario, email: e.target.value })
                }
                required
              />
              <input
                type="password"
                placeholder="Senha"
                className="w-full p-2 rounded bg-slate-700 focus:outline-none"
                value={novoUsuario.senha}
                onChange={(e) =>
                  setNovoUsuario({ ...novoUsuario, senha: e.target.value })
                }
                required
              />
              <select
                className="w-full p-2 rounded bg-slate-700 focus:outline-none"
                value={novoUsuario.tipo}
                onChange={(e) =>
                  setNovoUsuario({ ...novoUsuario, tipo: e.target.value })
                }
              >
                <option value="cliente">Cliente</option>
                <option value="admin">Admin</option>
              </select>

              <div className="flex justify-end gap-2 mt-4">
                <button
                  type="button"
                  onClick={() => setShowCreate(false)}
                  className="bg-slate-600 hover:bg-slate-500 px-4 py-2 rounded font-semibold"
                >
                  Cancelar
                </button>
                <button
                  type="submit"
                  className="bg-green-600 hover:bg-green-500 px-4 py-2 rounded font-semibold"
                >
                  Criar
                </button>
              </div>
            </form>
          </div>
        </div>
      )}

      {/* Modal Editar */}
      {showEdit && (
        <div className="fixed inset-0 bg-black/60 flex items-center justify-center z-50">
          <div className="bg-slate-800 p-6 rounded-lg shadow-lg w-[420px]">
            <h2 className="text-xl font-semibold mb-4">Editar Usuário</h2>
            <form onSubmit={handleEdit} className="space-y-3">
              <input
                type="text"
                placeholder="Nome"
                className="w-full p-2 rounded bg-slate-700 focus:outline-none"
                value={formEdicao.nome}
                onChange={(e) =>
                  setFormEdicao({ ...formEdicao, nome: e.target.value })
                }
                required
              />
              <input
                type="email"
                placeholder="Email"
                className="w-full p-2 rounded bg-slate-700 focus:outline-none"
                value={formEdicao.email}
                onChange={(e) =>
                  setFormEdicao({ ...formEdicao, email: e.target.value })
                }
                required
              />
              <select
                className="w-full p-2 rounded bg-slate-700 focus:outline-none"
                value={formEdicao.tipo}
                onChange={(e) =>
                  setFormEdicao({ ...formEdicao, tipo: e.target.value })
                }
              >
                <option value="cliente">Cliente</option>
                <option value="admin">Admin</option>
              </select>

              {/* senha opcional */}
              <input
                type="password"
                placeholder="Nova senha (opcional)"
                className="w-full p-2 rounded bg-slate-700 focus:outline-none"
                value={formEdicao.senha}
                onChange={(e) =>
                  setFormEdicao({ ...formEdicao, senha: e.target.value })
                }
              />

              <div className="flex justify-end gap-2 mt-4">
                <button
                  type="button"
                  onClick={() => setShowEdit(false)}
                  className="bg-slate-600 hover:bg-slate-500 px-4 py-2 rounded font-semibold"
                >
                  Cancelar
                </button>
                <button
                  type="submit"
                  className="bg-yellow-500 hover:bg-yellow-400 text-black px-4 py-2 rounded font-semibold"
                >
                  Salvar
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}
