FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o 1337b04rd ./cmd/1337b04rd

EXPOSE ${REST_API_PORT}
CMD [ "./1337d0rd" ]