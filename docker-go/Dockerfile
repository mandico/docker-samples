# Etapa de compilação
FROM golang:latest AS builder

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar os arquivos do módulo Go e baixar as dependências
COPY go.mod ./
RUN go mod download

# Copiar o código fonte para o diretório de trabalho
COPY . .

# Compilar o aplicativo para um binário estático
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Etapa final, usando uma imagem base mínima
FROM scratch

# Copiar o binário compilado da etapa de compilação
COPY --from=builder /app/main .

# Definir o comando para executar o binário
CMD ["./main"]