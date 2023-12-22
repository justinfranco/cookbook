FROM golang:1.21 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /cookbook-backend

FROM alpine:latest

COPY --from=build /cookbook-backend /

EXPOSE 3000

CMD ["/cookbook-backend"]