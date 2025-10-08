import type { JSX } from "react";
import { Navigate } from "react-router-dom";

interface ProtectedRouteProps {
  children: JSX.Element;
}

export default function ProtectedRoute({ children }: ProtectedRouteProps) {
  const isAuthenticated = localStorage.getItem("token"); // simulação

  if (!isAuthenticated) {
    return <Navigate to="/" replace />;
  }

  return children;
}
