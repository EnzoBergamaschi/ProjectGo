import { jwtDecode } from "jwt-decode";

export interface DecodedToken {
  id: number;
  email: string;
  tipo: string;
  exp: number;
  iat: number;
}

export function decodeToken(token: string): DecodedToken | null {
  try {
    return jwtDecode<DecodedToken>(token);
  } catch (error) {
    console.error("Erro ao decodificar token:", error);
    return null;
  }
}

export function isTokenExpired(token: string): boolean {
  try {
    const decoded = jwtDecode<DecodedToken>(token);
    const currentTime = Date.now() / 1000;
    return decoded.exp < currentTime;
  } catch {
    return true;
  }
}
