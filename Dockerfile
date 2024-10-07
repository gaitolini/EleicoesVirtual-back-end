# Etapa 1: Imagem base para build
FROM golang:1.20 AS builder

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar o go.mod e go.sum e baixar dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante dos arquivos do projeto
COPY . .

# Compilar o binário
RUN go build -o eleicoesvirtual

# Etapa 2: Imagem final
FROM gcr.io/distroless/base-debian10

# Copiar o binário da etapa de build
COPY --from=builder /app/eleicoesvirtual /eleicoesvirtual

# Expor a porta utilizada pelo serviço
EXPOSE 8080

# Comando de execução do container
ENTRYPOINT ["/eleicoesvirtual"]
