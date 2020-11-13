FROM golang:alpine

WORKDIR /app
COPY . .
RUN go get -v

CMD ["app"]
