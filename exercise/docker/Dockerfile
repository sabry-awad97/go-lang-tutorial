FROM golang:alpine
RUN mkdir /files
COPY main.go /files
WORKDIR /files
RUN go build -o /files/main main.go
ENTRYPOINT ["/files/main"]
