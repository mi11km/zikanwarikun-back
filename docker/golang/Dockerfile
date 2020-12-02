# 本番環境用
FROM golang:1.14-alpine as builder
ENV GO111MODULE=on
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/main -ldflags '-s -w'

FROM scratch as runner
EXPOSE 8080
COPY --from=builder /go/bin/main /app/main
CMD ["/app/main"]