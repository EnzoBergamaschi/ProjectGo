import { useState } from "react";
import { registerUser } from "../services/userService";

export default function RegisterPage() {
  const [nome, setNome] = useState("");
  const [email, setEmail] = useState("");
  const [senha, setSenha] = useState("");
  const [confirmar, setConfirmar] = useState("");
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError("");
    setSuccess("");

    if (senha !== confirmar) {
      setError("As senhas não coincidem.");
      return;
    }

    try {
      await registerUser({ nome, email, senha });
      setSuccess("Usuário registrado com sucesso! Agora você pode fazer login.");
      setNome("");
      setEmail("");
      setSenha("");
      setConfirmar("");
    } catch (err: any) {
      console.error(err);
      setError("Erro ao registrar usuário. Tente novamente.");
    }
  }

  return (
    <div className="flex min-h-screen items-center justify-center bg-slate-900 text-white">
      <div className="w-96">
        <form
          onSubmit={handleSubmit}
          className="bg-slate-800 p-8 rounded-2xl shadow-md space-y-4"
        >
          <h2 className="text-2xl font-bold text-center">Criar conta</h2>

          <input
            type="text"
            placeholder="Nome completo"
            value={nome}
            onChange={(e) => setNome(e.target.value)}
            className="w-full p-2 rounded bg-slate-700 focus:outline-none"
            required
          />

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

          <input
            type="password"
            placeholder="Confirmar senha"
            value={confirmar}
            onChange={(e) => setConfirmar(e.target.value)}
            className="w-full p-2 rounded bg-slate-700 focus:outline-none"
            required
          />

          {error && <p className="text-red-400 text-sm">{error}</p>}
          {success && <p className="text-green-400 text-sm">{success}</p>}

          <button
            type="submit"
            className="w-full bg-blue-600 hover:bg-blue-500 transition p-2 rounded font-semibold"
          >
            Registrar
          </button>
        </form>

        <p className="text-center text-sm text-gray-400 mt-4">
          Já tem uma conta?{" "}
          <a href="/" className="text-blue-400 hover:underline">
            Faça login
          </a>
        </p>
      </div>
    </div>
  );
}
