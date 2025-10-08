import { Link } from "react-router-dom";

export default function Sidebar() {
  return (
    <aside className="w-64 bg-slate-900 text-white h-screen p-4">
      <ul className="space-y-4">
        <li>
          <Link to="/dashboard" className="hover:text-blue-400">
            Dashboard
          </Link>
        </li>
        <li>
          <Link to="/produtos" className="hover:text-blue-400">
            Produtos
          </Link>
        </li>
        <li>
          <Link to="/vendas" className="hover:text-blue-400">
            Vendas
          </Link>
        </li>
      </ul>
    </aside>
  );
}
