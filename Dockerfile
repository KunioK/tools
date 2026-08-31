FROM golang:1.23-alpine AS builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
COPY ./tools_service/go.mod ./
RUN go mod download
COPY tools_service/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o tools-service .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/tools-service .
EXPOSE 8080
CMD ["./tools-service"]
