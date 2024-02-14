FROM golang

WORKDIR /app

COPY . /app/

RUN go mod tidy
RUN go build cmd/main.go

CMD ["./main"]