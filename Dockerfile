# Usamos imagen oficial de Go
FROM golang:1.23

# Creamos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos los archivos del proyecto al contenedor
COPY . .

# Descargamos las dependencias
RUN go mod download

# Compilamos el binario
RUN go build -o delete_product

# Exponemos el puerto en el que corre el servicio (ajústalo si es diferente)
EXPOSE 4004

# Comando para iniciar el microservicio
CMD ["./delete_product"]