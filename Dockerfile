FROM golang:1.22
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o myImage
ENTRYPOINT ["/app/myImage"]