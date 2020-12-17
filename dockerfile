FROM golang
RUN go build wiki.go
EXPOSE 8081
RUN go run wiki.go 8081

