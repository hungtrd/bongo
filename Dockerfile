FROM golang:1.19-alpine

WORKDIR /bongo

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN GOARCH=arm GOOS=linux go build -o main .

EXPOSE 80801

CMD ["./main"]
