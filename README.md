ProjetoGo - Go (backend) + React (frontend) - Enzo Teixeira Bergamaschi  

Este foi meu primeiro readme,desculpa qualquer dor de cabeça.

1. Clonar o repositório

git clone https://github.com/EnzoBergamaschi/ProjectGo.git
cd ProjectGo

2. Construir e subir os containers do backend (API + MySQL)

docker compose up -d

docker compose build api

docker compose up --build api

3. Aplicar as migrations (criar as tabelas do banco)

Após o MySQL estar rodando, entre no container do banco:

docker exec -it projectgo-db mysql -u root -prootpass projectgo

Agora, copie e cole o conteúdo do arquivo SQL localizado em: #Tentei deixar automatico mas nao deu.

Api/migrations/001_migrate.sql

lembrando que neste migrate 

-- Cria o administrador padrão (senha: 1234) (admin@email.com)

4. Rodar o frontend (Web)

Entre na pasta Web:
cd Web

Instale as dependências:
npm install

inicie o servidor de desenvolvimento:
npm run dev

frontend estará disponível em:
http://localhost:5173/