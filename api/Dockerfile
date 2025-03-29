FROM golang:1.24.1-alpine3.21

WORKDIR /app/

COPY . /app/

RUN go mod download

CMD [ "go", "run", "main.go" ]