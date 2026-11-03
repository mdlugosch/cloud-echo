# Stufe 1: Kompilierung
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server main.go
# Stufe 2: Schlankes Laufzeit-Image
FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["./server"]
