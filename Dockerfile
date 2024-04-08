FROM golang:1.23
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o myImage
ENTRYPOINT ["/app/myImage"]
