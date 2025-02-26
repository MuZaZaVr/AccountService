FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o account-service cmd/main.go

CMD ["./account-service"]