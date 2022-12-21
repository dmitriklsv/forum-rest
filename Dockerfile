FROM golang:1.19-alpine
WORKDIR /app

COPY . .

LABEL version="1.0" 
LABEL creator="@arturzhamaliyev"
RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init
RUN go mod download
RUN go build cmd/main.go
EXPOSE 8080

CMD ["./main"]