interface NavbarProps {
  onLogout: () => void;
}

export default function Navbar({ onLogout }: NavbarProps) {
  return (
    <nav className="bg-slate-800 shadow-md py-4 px-6 flex justify-between items-center">
      <h1 className="text-2xl font-bold text-blue-400">ProjectGo</h1>
      <button
        onClick={onLogout}
        className="bg-red-600 hover:bg-red-500 transition px-4 py-2 rounded font-semibold"
      >
        Sair
      </button>
    </nav>
  );
}
