FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# RUN go build -o 1337b04rd ./cmd/main.go

EXPOSE ${REST_API_PORT}
CMD [ "./1337b04rd" ]