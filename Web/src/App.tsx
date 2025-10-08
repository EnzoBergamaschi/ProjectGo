import { BrowserRouter, Routes, Route } from "react-router-dom";
import LoginPage from "./pages/LoginPage";
import RegisterPage from "./pages/RegisterPage";
import DashboardPage from "./pages/Dashboard";
import ProdutosPage from "./pages/ProdutosPage";
import VendasPage from "./pages/VendasPage";
import UsuariosPage from "./pages/UsuariosPage";

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/dashboard" element={<DashboardPage />} />
        <Route path="/produtos" element={<ProdutosPage />} />
        <Route path="/vendas" element={<VendasPage />} />
        <Route path="/usuarios" element={<UsuariosPage />} />
      </Routes>
    </BrowserRouter>
  );
}
