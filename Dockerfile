# Etapa de build
FROM golang:1.21 AS build

WORKDIR /app

# Copiar o código para o contêiner
COPY . .

# Rodar go mod tidy e compilar o binário
RUN go mod tidy && go build -o eleicoes-backend ./main.go

# Etapa final para criar uma imagem leve
FROM alpine:latest
WORKDIR /root/

# Instalar ca-certificates para chamadas HTTPS do Go
RUN apk --no-cache add ca-certificates

# Copiar o binário gerado para a imagem final
COPY --from=build /app/eleicoes-backend .

# Definir permissões de execução para o binário
RUN chmod +x ./eleicoes-backend

# Comando para rodar o binário
CMD ["./eleicoes-backend"]
