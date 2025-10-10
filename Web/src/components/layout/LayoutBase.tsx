import { Outlet } from "react-router-dom";
import Navbar from "./Navbar";
import Sidebar from "./Sidebar";

export default function LayoutBase() {
  return (
    <div className="flex h-screen bg-slate-800 text-white">
      <Sidebar />
      <div className="flex flex-col flex-1">
        <Navbar onLogout={() => {}} />
        <main className="p-6 flex-1 overflow-y-auto">
          <Outlet /> 
        </main>
      </div>
    </div>
  );
}
