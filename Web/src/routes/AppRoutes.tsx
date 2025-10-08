import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "../pages/LoginPage";
import Dashboard from "../pages/Dashboard";
import LayoutBase from "../components/layout/LayoutBase";
import ProtectedRoute from "../components/ProtectedRoute";
import UsuariosPage from "../pages/UsuariosPage";
import ProdutosPage from "../pages/ProdutosPage";
import VendasPage from "../pages/VendasPage";

export default function AppRoutes() {
  return (
    <BrowserRouter>
      <Routes>
        {/* Rota pública */}
        <Route path="/" element={<Login />} />

        {/* Rota protegida com layout */}
        <Route
          element={
            <ProtectedRoute>
              <LayoutBase />
            </ProtectedRoute>
          }
        >
          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/usuarios" element={<UsuariosPage />} />
          <Route path="/produtos" element={<ProdutosPage />} />
          <Route path="/vendas" element={<VendasPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}
