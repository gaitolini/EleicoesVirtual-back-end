# Etapa de build
FROM golang:1.22 AS build

WORKDIR /app

# Copiar o código para o contêiner
COPY . .

# Rodar go mod tidy e build com compilação estática
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o eleicoes-backend ./main.go

# Etapa final para criar uma imagem leve
FROM alpine:latest
WORKDIR /app/

# Instalar ca-certificates, necessário para chamadas HTTPS do Go
RUN apk --no-cache add ca-certificates

# Copiar o binário gerado para a imagem final
COPY --from=build /app/eleicoes-backend .

# Copiar o arquivo de credenciais
COPY eleicoesvirtual-firebase-adminsdk-baotz-3973687bb4.json /app/credentials/

# Definir a variável de ambiente para as credenciais do Firebase
ENV GOOGLE_APPLICATION_CREDENTIALS=/app/credentials/eleicoesvirtual-firebase-adminsdk-baotz-3973687bb4.json

# Comando para rodar o binário
CMD ["./eleicoes-backend"]
