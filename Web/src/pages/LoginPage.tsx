import { useState } from "react";
import { loginUser } from "../services/authService";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [senha, setSenha] = useState("");
  const [error, setError] = useState("");

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError("");

    try {
      const res = await loginUser({ email, senha });
      localStorage.setItem("token", res.token);
      console.log("Login bem-sucedido:", res.token);
      alert("Login realizado com sucesso!");
      window.location.href = "/dashboard";
    } catch (err: any) {
      console.error(err);
      setError("Credenciais inválidas ou erro no servidor.");
    }
  }

  return (
    <div className="flex min-h-screen items-center justify-center bg-slate-900 text-white">
      <div className="w-96">
        <form
          onSubmit={handleSubmit}
          className="bg-slate-800 p-8 rounded-2xl shadow-md space-y-4"
        >
          <h2 className="text-2xl font-bold text-center">Login</h2>

          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="w-full p-2 rounded bg-slate-700 focus:outline-none"
            required
          />

          <input
            type="password"
            placeholder="Senha"
            value={senha}
            onChange={(e) => setSenha(e.target.value)}
            className="w-full p-2 rounded bg-slate-700 focus:outline-none"
            required
          />

          {error && <p className="text-red-400 text-sm">{error}</p>}

          <button
            type="submit"
            className="w-full bg-blue-600 hover:bg-blue-500 transition p-2 rounded font-semibold"
          >
            Entrar
          </button>
        </form>

        <p className="text-center text-sm text-gray-400 mt-4">
          Ainda não tem uma conta?{" "}
          <a href="/register" className="text-blue-400 hover:underline">
            Cadastre-se
          </a>
        </p>
      </div>
    </div>
  );
}
