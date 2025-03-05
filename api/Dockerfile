FROM golang:1.24.1-alpine3.21

WORKDIR /app/

COPY go.mod go.sum /app/

RUN go mod download

COPY . /app/

CMD [ "go", "run", "main.go" ]