FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/myapp .

RUN mkdir -p /app/logs

EXPOSE 8086

CMD ["./myapp"]
