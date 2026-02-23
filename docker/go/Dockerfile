FROM golang:1.22-alpine

WORKDIR /app

# install git in linux alpine for go mod tidy
RUN apk add --no-cache git

# install air pkg for hot reload
RUN go install github.com/air-verse/air@latest

EXPOSE 8080

CMD ["air"]