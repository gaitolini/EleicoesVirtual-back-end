# Etapa inicial para construir o contêiner
FROM golang:1.19 AS build

WORKDIR /app

# Copiar o código para o contêiner
COPY . .

# Alterar a versão do Go para 1.19 no go.mod e remover a diretiva toolchain
RUN sed -i 's/go 1\.21/go 1\.19/' go.mod && sed -i '/toolchain/d' go.mod

# Baixar as dependências e compilar
RUN go mod tidy && go build -o eleicoes-backend

# Etapa final para criar uma imagem leve
FROM alpine:latest

# Configurar a variável de ambiente para o Firebase
ENV GOOGLE_APPLICATION_CREDENTIALS=/root/eleicoesvirtual-firebase-adminsdk-baotz-3973687bb4.json

WORKDIR /root/
COPY --from=build /app/eleicoes-backend .
CMD ["./eleicoes-backend"]
