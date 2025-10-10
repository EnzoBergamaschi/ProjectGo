-- ===========================================
-- 001_init.sql — Estrutura completa do banco ProjectGo
-- ===========================================

-- ========== TABELA DE USUÁRIOS ==========
CREATE TABLE IF NOT EXISTS usuarios (
  id INT AUTO_INCREMENT PRIMARY KEY,
  nome VARCHAR(100) NOT NULL,
  email VARCHAR(150) NOT NULL UNIQUE,
  senha_hash VARCHAR(255) NOT NULL,
  tipo ENUM('admin','cliente') NOT NULL DEFAULT 'cliente',
  criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
    
-- ========== TABELA DE PRODUTOS ==========
CREATE TABLE IF NOT EXISTS produtos (
  id INT AUTO_INCREMENT PRIMARY KEY,
  nome VARCHAR(100) NOT NULL,
  descricao TEXT,
  preco DECIMAL(10,2) NOT NULL DEFAULT 0.00,
  estoque INT NOT NULL DEFAULT 0,
  criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ========== TABELA DE VENDAS ==========
CREATE TABLE IF NOT EXISTS vendas (
  id INT AUTO_INCREMENT PRIMARY KEY,
  id_usuario INT NOT NULL,
  total DECIMAL(10,2) NOT NULL DEFAULT 0.00,
  status ENUM('pendente','pago','enviado','cancelado') NOT NULL DEFAULT 'pendente',
  criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (id_usuario) REFERENCES usuarios(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- ========== TABELA DE ITENS DE VENDA ==========
CREATE TABLE IF NOT EXISTS itens_venda (
  id INT AUTO_INCREMENT PRIMARY KEY,
  id_venda INT NOT NULL,
  id_produto INT NOT NULL,
  quantidade INT NOT NULL,
  preco_unitario DECIMAL(10,2) NOT NULL,
  FOREIGN KEY (id_venda) REFERENCES vendas(id) ON DELETE CASCADE,
  FOREIGN KEY (id_produto) REFERENCES produtos(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ========== DADOS INICIAIS ==========
-- Cria o administrador padrão (senha: 1234)
INSERT INTO usuarios (nome, email, senha_hash, tipo)
VALUES (
  'Admin',
  'admin@email.com',
  '$2b$12$0S3EhR8FbRDEGr4jAq/Vl.GX7PIdP6aPPkEx.bcJt.h9T5ZT5aV9.',
  'admin'
)
ON DUPLICATE KEY UPDATE tipo = VALUES(tipo);

-- Produtos de exemplo
INSERT INTO produtos (nome, descricao, preco, estoque)
VALUES
('Camisa Milan', 'Camisa oficial do Milan - temporada 2025', 249.90, 10),
('Bola de Futebol', 'Bola oficial de campo da Nike', 159.90, 20),
('Chuteira Predator', 'Chuteira Adidas Predator FG', 499.00, 5);

-- Venda de exemplo (para teste de relacionamento)
INSERT INTO vendas (id_usuario, total, status)
VALUES (1, 0, 'aberta');

-- Itens da venda de exemplo
INSERT INTO itens_venda (id_venda, id_produto, quantidade, preco_unitario)
VALUES (1, 1, 2, 249.90);
