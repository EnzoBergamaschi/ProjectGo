import axios from "axios";

// Configura a URL base da sua API Go (container dev exposto em 8080)
const api = axios.create({
  baseURL: "http://localhost:8080",
  headers: {
    "Content-Type": "application/json",
  },
});

// Intercepta todas as requisições e inclui o token JWT, se existir
api.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token && config.headers) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default api;
