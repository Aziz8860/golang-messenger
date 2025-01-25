FROM golang:1.22.2-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

# copy semua file ke dalam container
COPY . .

RUN go build -o simple-messaging-app

# ngasih akses execute
RUN chmod +x simple-messaging-app

EXPOSE 4000

EXPOSE 8080

CMD ["./simple-messaging-app"]
