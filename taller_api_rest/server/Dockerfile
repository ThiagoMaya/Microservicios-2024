FROM golang:latest

# Copiar el código de la aplicación
COPY . /app

# Establecer el directorio de trabajo
WORKDIR /app

# Instalar las dependencias de la aplicación
RUN go mod tidy

# Compilar la aplicación
RUN go build -o main .

# Ejecutar la aplicación
CMD ["./main"]