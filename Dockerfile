# Usar uma imagem base do Go para compilar o código
FROM golang:1.19 as build

WORKDIR /app

# Copiar os arquivos de código
COPY . .

# Construir o binário de forma estática
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o eleicoes-backend

# Usar uma imagem base mais leve para executar o binário
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar o binário do estágio anterior
COPY --from=build /app/eleicoes-backend .

# Expor a porta 8080
EXPOSE 8080

# Comando para iniciar a aplicação
CMD ["./eleicoes-backend"]
