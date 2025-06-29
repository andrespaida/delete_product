FROM golang:1.23
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o delete_product
EXPOSE 4004
CMD ["./delete_product"]