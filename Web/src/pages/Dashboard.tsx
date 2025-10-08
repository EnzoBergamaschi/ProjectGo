import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Navbar from "../components/layout/Navbar";
import { decodeToken, isTokenExpired } from "../utils/jwtHelper";
import type { DecodedToken } from "../utils/jwtHelper";

export default function DashboardPage() {
  const [user, setUser] = useState<DecodedToken | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      navigate("/");
      return;
    }

    if (isTokenExpired(token)) {
      alert("Sessão expirada. Faça login novamente.");
      localStorage.removeItem("token");
      navigate("/");
      return;
    }

    const decoded = decodeToken(token);
    if (decoded) {
      setUser(decoded);
    } else {
      navigate("/");
    }
  }, [navigate]);

  function handleLogout() {
    localStorage.removeItem("token");
    navigate("/");
  }

  return (
    <div className="flex flex-col min-h-screen bg-slate-900 text-white">
      <Navbar onLogout={handleLogout} />

      <main className="flex-1 p-8">
        <h1 className="text-3xl font-bold mb-4">Bem-vindo ao ProjectGo</h1>

        {user && (
          <p className="text-gray-300 mb-8">
            Olá, <span className="font-semibold">{user.email}</span> ({user.tipo})
          </p>
        )}

        <div className="grid grid-cols-1 sm:grid-cols-3 gap-6">
          {/* Card Produtos */}
          <div className="bg-slate-800 p-6 rounded-xl shadow-md hover:shadow-lg transition">
            <h2 className="text-xl font-semibold mb-2">Produtos</h2>
            <p className="text-gray-400 mb-3">Gerencie seu estoque de produtos.</p>
            <button
              onClick={() => navigate("/produtos")}
              className="bg-blue-600 hover:bg-blue-500 transition px-4 py-2 rounded font-semibold"
            >
              Acessar
            </button>
          </div>

          {/* Card Vendas */}
          <div className="bg-slate-800 p-6 rounded-xl shadow-md hover:shadow-lg transition">
            <h2 className="text-xl font-semibold mb-2">Vendas</h2>
            <p className="text-gray-400 mb-3">Registre e visualize suas vendas.</p>
            <button
              onClick={() => navigate("/vendas")}
              className="bg-green-600 hover:bg-green-500 transition px-4 py-2 rounded font-semibold"
            >
              Acessar
            </button>
          </div>

          {/* Card Usuários */}
          <div className="bg-slate-800 p-6 rounded-xl shadow-md hover:shadow-lg transition">
            <h2 className="text-xl font-semibold mb-2">Usuários</h2>
            <p className="text-gray-400 mb-3">Administre usuários do sistema.</p>
            <button
              onClick={() => navigate("/usuarios")}
              className="bg-purple-600 hover:bg-purple-500 transition px-4 py-2 rounded font-semibold"
            >
              Acessar
            </button>
          </div>
        </div>
      </main>
    </div>
  );
}
