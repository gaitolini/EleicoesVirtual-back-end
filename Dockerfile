# Etapa de build
FROM golang:1.22 AS build

WORKDIR /app

# Copiar o código para o contêiner
COPY . .

# Rodar go mod tidy e build com compilação estática
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go mod tidy && go build -o eleicoes-backend ./main.go

# Etapa final para criar uma imagem leve
FROM alpine:latest
WORKDIR /app/

# Instalar ca-certificates, necessário para chamadas HTTPS do Go
RUN apk --no-cache add ca-certificates

# Criar o diretório para as credenciais
RUN mkdir -p /app/credentials

# Copiar o binário gerado para a imagem final
COPY --from=build /app/eleicoes-backend .

# Definir a variável de ambiente para as credenciais do Firebase
ENV GOOGLE_APPLICATION_CREDENTIALS=/app/credentials/eleicoesvirtual-firebase-adminsdk-baotz-3973687bb4.json

# Definir a porta que a aplicação vai escutar (opcional, para documentação)
EXPOSE 8081

# Comando para rodar o binário
CMD ["./eleicoes-backend"]
