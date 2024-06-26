# focus-finance

Esta é uma API restfull desenvolvida para que possa salvar e consultar gastos e recebimentos, podendo ser feita a separação por meses.

## Tecnologias Utilizadas
- [Golang]: linguagem de programação criada pela Google. 
- [GIN]: Web framework.
- [PostgreSQL]: Banco de dados relacional.
- [Arquitetura MVC] para melhorar a organização do projeto implementei a arquitetura MVC - model, view e controller -  com camadas de repository para operações com banco de dados e service para separar a regra de negócio.
- [GORM] ORM para ajudar/agilizar no momento da criação de consultas e criaçao de tabelas SQL.

## Instalação
- Clone o repositório
- navegue até o diretorio do projeto
- execute "go get" para baixar todas as dependencias
- rode o projeto com "go run focus-finance"
- acesse http://localhost:5000
