# Usando a imagem oficial do Golang para criar a build
FROM golang:1.19 AS build

# Configurar o diretório de trabalho dentro do container
WORKDIR /app

# Copiar os arquivos do projeto para dentro do container
COPY . .

# Baixar as dependências e construir o binário
RUN go mod tidy && go build -o eleicoes-backend

# Criar uma imagem mínima para rodar o binário
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar o binário construído na etapa anterior
COPY --from=build /app/eleicoes-backend .

# Expor a porta 8080 para acesso externo
EXPOSE 8080

# Comando para executar o serviço
CMD ["./eleicoes-backend"]
