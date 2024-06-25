FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o openfeature-sample -a -ldflags '-extldflags "-static"'

FROM gcr.io/distroless/static


WORKDIR /app/
COPY --from=builder /app/openfeature-sample .
EXPOSE 8888

CMD ["./openfeature-sample"]
