FROM golang:1.17
RUN mkdir /app
ADD ./app /app
WORKDIR /app
RUN go mod download
RUN go build -o app .
CMD ["/app/app"]