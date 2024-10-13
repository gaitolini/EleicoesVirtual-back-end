# Etapa de build para compilar o projeto Go
FROM golang:1.22 AS build

# Definir o diretório de trabalho
WORKDIR /app

# Copiar o go.mod e go.sum para evitar refazer o download das dependências em cada build
COPY go.mod go.sum ./

# Rodar go mod download para baixar as dependências
RUN go mod download

# Copiar o restante do código para o contêiner
COPY . .

# Compilar o projeto com compilação estática
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o eleicoes-backend ./main.go

# Etapa final para criar uma imagem leve
FROM alpine:latest

# Definir o diretório de trabalho
WORKDIR /app

# Instalar somente o necessário para chamadas HTTPS do Go
RUN apk --no-cache add ca-certificates

# Copiar o binário gerado na etapa de build para a imagem final
COPY --from=build /app/eleicoes-backend .

# Copiar o arquivo de credenciais
COPY firebaseConfig.json /app/credentials/firebaseConfig.json

# Definir a variável de ambiente para as credenciais do Firebase
ENV GOOGLE_APPLICATION_CREDENTIALS=/app/credentials/firebaseConfig.json

# Limpar o cache do apk para economizar espaço (otimização extra)
RUN rm -rf /var/cache/apk/*

# Comando para rodar o binário
CMD ["./eleicoes-backend"]
