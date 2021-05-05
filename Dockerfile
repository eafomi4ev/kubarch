FROM golang:1.16.3 as builder
WORKDIR /app
COPY . .

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container

# Download all the dependencies
#RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./...


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["./main"]