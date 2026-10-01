<<<<<<< HEAD

FROM golang:1.22-alpine AS builder

WORKDIR /app


COPY go.mod ./
RUN go mod download || true

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortener main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/url-shortener .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

EXPOSE 8080

=======

FROM golang:1.22-alpine AS builder

WORKDIR /app


COPY go.mod ./
RUN go mod download || true

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortener main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/url-shortener .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

EXPOSE 8080

>>>>>>> e14adc1597877f2a12868b8a89de88ee0786570e
CMD ["./url-shortener"]