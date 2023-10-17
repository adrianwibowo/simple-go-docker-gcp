FROM golang:1.20-alpine

WORKDIR /app
COPY . .

RUN go build -o simpe-go-docker-gcp-app

EXPOSE 9000

CMD ./simpe-go-docker-gcp-app
