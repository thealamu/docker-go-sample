FROM golang:alpine

WORKDIR /app
COPY . .
RUN go get -v
RUN go build -o docker-go-sample

CMD ["./docker-go-sample"]
