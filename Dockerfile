# Usar a imagem base do Go
FROM golang:1.19 AS build

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar o código do projeto para o contêiner
COPY . .

# Baixar as dependências e compilar o binário
RUN go mod tidy && go build -o eleicoes-backend

# Criar uma imagem mínima para rodar o binário
FROM alpine:latest

# Configurar a variável de ambiente para o Firebase
ENV GOOGLE_APPLICATION_CREDENTIALS=/root/eleicoesvirtual-firebase-adminsdk-baotz-3973687bb4.json

# Copiar o binário e o arquivo de credenciais do Firebase para o contêiner
COPY --from=build /app/eleicoes-backend /root/
COPY --from=build /app/eleicoesvirtual-firebase-adminsdk-baotz-3973687bb4.json /root/

# Expor a porta em que o servidor irá rodar
EXPOSE 8080

# Comando para rodar o binário
CMD ["/root/eleicoes-backend"]
