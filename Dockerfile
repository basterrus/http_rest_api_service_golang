FROM golang:1.17

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgres-client

RUN go mod download
RUN go build -o todo-app ./cmd/main.go

CMD ["./todo-app"]