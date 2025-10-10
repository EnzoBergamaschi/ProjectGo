ProjetoGo - Go (backend) + React (frontend) - Enzo Teixeira Bergamaschi  

Este foi meu primeiro readme,desculpa qualquer dor de cabeça.

1. Clonar o repositório

git clone https://github.com/EnzoBergamaschi/ProjectGo.git

cd ProjectGo

2. Construir e subir os containers do backend (API + MySQL)

Na raiz do projeto, execute:
> docker compose up --build api

Esses comandos criam e sobem os containers do MySQL e da API Go, montando o banco de dados e iniciando a API automaticamente.

3. Rodar o frontend (Web)

Entre na pasta Web:
cd Web

Instale as dependências:
npm install

inicie o servidor de desenvolvimento:
npm run dev

frontend estará disponível em:
http://localhost:5173/