# Eleições Virtuais - Backend 🌟

Bem-vindo ao repositório do Backend da aplicação Eleições Virtuais! Este projeto foi desenvolvido para ser um simulador de eleição virtual, promovendo o valor da democracia e a importância do processo eleitoral. É uma parte do meu portfólio e uma demonstração das minhas habilidades em desenvolvimento backend, infraestrutura em cloud, automação de pipelines CI/CD e gestão de projetos de software.

Link para o frontend: [Eleições Virtuais - Frontend](https://github.com/gaitolini/EleicoesVirtual-front-end)

Link do projeto online: [Eleições Virtuais](https://eleicoesvirtual.web.app/)

## 🚀 Tecnologias e Ferramentas Utilizadas

- Go Lang 🛠️
- GitHub & GitHub Actions 🛢️
- Docker 🛠️
- Firebase 🔖
- AWS EC2 🏠
- Cloudflare Zero Trust ⚡️
- Linux 🛡️
- VSCode 🔧
- Node.js & npm 🤖

## 🛍️ O Que Foi Feito

### 💼 Criação e Configuração da Infraestrutura

1. **Criação do Repositório no GitHub**: Repositório configurado para armazenamento do código e controle de versão.
2. **Instância EC2 na AWS**: Criei uma instância de EC2 para hospedar o backend da aplicação com sistema operacional Linux e adicionei um volume de 16GB para armazenamento de dados.
3. **Configuração da EC2**:
   - **Instalação do Docker** para executar contêineres do backend.
   - **Configuração do Git** e criação de chave SSH para acessar o repositório remoto.

### 🛠️ Configuração do Cloudflare Zero Trust

- **Criação do DNS** personalizado com domínio `api.gaitolini.com.br` para acessar a API.
- **Configuração do Tunnel Zero Trust** para garantir conexão segura entre a aplicação e o público.

### ⚙️ Pipeline CI/CD no GitHub Actions

- **Workflow do CI/CD** para automação das etapas de build e deploy da aplicação.
  ```yaml
  .github/workflows/deploy-backend.yml
  ```
- **Dockerfile** para criar a imagem do backend.
  ```Dockerfile
  FROM golang:1.18
  COPY . /app
  RUN go build -o eleicoes-backend
  CMD ["./eleicoes-backend"]
  ```
- **Secrets Configurados** no GitHub, como chaves do Firebase, SSH e outras credenciais.

### 🛠️ Criação do Projeto no Firebase

- **Firebase Firestore**: Configurado para armazenar as eleições e seus respectivos dados.
- **FirebaseConfig.json** para autenticação e integração com o backend.

### 🚀 Backend em Go Lang

- **Camada Middleware de Autenticação CORS**: Implementada no projeto para garantir a segurança das requisições, validando os domínios de origem.
  ```go
  func CorsMiddleware(next http.Handler) http.Handler { /*...*/ }
  ```
- **Rotas de API Restful**: CRUD utilizando o pacote "mux" para lidar com as requisições HTTP.
  - **POST**: Criar Eleições
  - **GET**: Listar Eleições
  - **PUT**: Atualizar Eleições
  - **DELETE**: Deletar Eleições
- **Padrão MVC**: Dividido entre Controladores, Serviços e Modelos, facilitando a manutenção e organização do código.
  - **Controller**: `controllers/eleicoes.go`
  - **Service**: `services/eleicoes.go`
  - **Model**: `models/eleicao.go`

### 🌐 Testes e Validação

- **Postman**: Utilizado para testar todos os endpoints do backend, tanto localmente (localhost:8080) quanto na instância AWS e DNS personalizado (`api.gaitolini.com.br/eleicoes`).

## 🛡️ Objetivo do Projeto

O projeto tem como objetivo ser um MVP de um simulador de eleição virtual, incentivando a educação democrática e a compreensão do processo eleitoral. Esta versão inclui apenas o backend, sendo todo o frontend desenvolvido em um repositório à parte, que pode ser encontrado [aqui](https://github.com/gaitolini/EleicoesVirtual-front-end).

Este projeto também demonstra habilidades em:

- **Desenvolvimento Backend** utilizando Go Lang.
- **DevOps**: Infraestrutura em cloud com AWS, automação com GitHub Actions, e conteinerização com Docker.
- **CI/CD**: Automação do ciclo de vida do software, desde build até deploy.

Estou em busca de oportunidades como desenvolvedor backend, fullstack ou na área de DevOps e infraestrutura em cloud. Caso queira conhecer mais sobre o projeto ou me oferecer uma oportunidade de colaboração, sinta-se à vontade para entrar em contato.

## Contact 📮

Se você quer saber mais sobre este projeto ou discutir colaborações futuras, sinta-se à vontade para me contatar:

- LinkedIn: [Anderson Gaitolini](https://www.linkedin.com/in/andersongaitolini/)
- WhatsApp: [Entre em contato](https://youtu.be/IGP38bz-K48?si=62Khct2-dAFR3qn5)

## Estrutura do Projeto

```
.github
.github\workflows
.github\workflows\deploy-backend.yml
controllers
controllers\eleicoes.go
middleware
middleware\cors.go
models
models\eleicao.go
services
services\eleicoes.go
services\firebase.go
services\firestore_wrapper.go
utils
utils\error_handler.go
Dockerfile
firebaseConfig.json
go.mod
go.sum
main.go
README.md
```

