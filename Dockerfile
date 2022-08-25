# COMMANDS
# docker build -t uwaterloocourses .
# docker run --rm --name=uwaterloocourses -p 8000:80 uwaterloocourses

# GOLANG
FROM golang:1.19 AS builder
WORKDIR /go/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# START THE BACKEND
FROM scratch
COPY --from=builder /main /app/
WORKDIR /app
CMD ["./main"]