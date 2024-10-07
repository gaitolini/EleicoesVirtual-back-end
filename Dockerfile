# Etapa de build
FROM golang:1.21 AS build

WORKDIR /app

# Copiar o código para o contêiner
COPY . .

# Copiar o arquivo de credenciais do caminho na VM
COPY /home/ec2-user/credentials/eleicoesvirtual-firebase-adminsdk-baotz-3973687bb4.json /app/

# Definir a variável de ambiente para o caminho dentro do contêiner
ENV GOOGLE_APPLICATION_CREDENTIALS="/app/eleicoesvirtual-firebase-adminsdk-baotz-3973687bb4.json"

# Rodar go mod tidy e build
RUN go mod tidy && go build -o eleicoes-backend

# Etapa final para criar uma imagem leve
FROM alpine:latest
WORKDIR /root/

# Instalar ca-certificates, necessário para chamadas HTTPS do Go
RUN apk --no-cache add ca-certificates

# Copiar o binário gerado para a imagem final
COPY --from=build /app/eleicoes-backend .

# Copiar o arquivo de credenciais para a imagem final
COPY --from=build /app/eleicoesvirtual-firebase-adminsdk-baotz-3973687bb4.json .

# Definir a variável de ambiente na imagem final
ENV GOOGLE_APPLICATION_CREDENTIALS="/root/eleicoesvirtual-firebase-adminsdk-baotz-3973687bb4.json"

# Comando para rodar o binário
CMD ["./eleicoes-backend"]
