# Etapa de build
FROM golang:1.21 AS build

WORKDIR /app

# Copiar o código para o contêiner
COPY . .

# Rodar go mod tidy e build
RUN go mod tidy && go build -o eleicoes-backend

# Etapa final para criar uma imagem leve
FROM alpine:latest
WORKDIR /root/

# Instalar ca-certificates, necessário para chamadas HTTPS do Go
RUN apk --no-cache add ca-certificates

# Copiar o binário gerado para a imagem final
COPY --from=build /app/eleicoes-backend .

# Configurar a variável de ambiente para as credenciais do Firebase
# Aqui você pode apontar para o local onde o arquivo de credenciais será montado em tempo de execução
ENV GOOGLE_APPLICATION_CREDENTIALS=/root/credentials/eleicoesvirtual-firebase-adminsdk-baotz-3973687bb4.json

# Copiar o arquivo de credenciais para a imagem (apenas no momento do build, caso precise)
# Caso queira copiar do seu ambiente local durante o build, faça assim:
# COPY /caminho-local/eleicoesvirtual-firebase-adminsdk-baotz-3973687bb4.json /root/credentials/

# Comando para rodar o binário
CMD ["./eleicoes-backend"]
