# COMMANDS
# docker build -t uwaterloocourses .
# docker run -d -p 80:80 uwaterloocourses

# NODE
FROM node:16-alpine as client_builder
COPY /frontend /node/app/
WORKDIR /node/app/frontend/
RUN npm install && npm run build

# GOLANG
FROM golang:latest as base_builder
WORKDIR /go/app
# copy go files
COPY go.mod ./
COPY go.sum ./
# go download
RUN go mod download
COPY *.go ./
# download server files
RUN go get github.com/realTristan/uwaterloo.courses
# run the server
RUN go build -o /server

# FOR SMALLER CONTAINERS
FROM gcr.io/distroless/base
COPY --from=base_builder /go/src/app/public /public/
COPY --from=base_builder /go/src/app/.env /
COPY --from=base_builder /go/bin/app /
COPY --from=client_builder /node/app/public/build/ /public/build/

# RUN APP
CMD ["/app", "/server"]