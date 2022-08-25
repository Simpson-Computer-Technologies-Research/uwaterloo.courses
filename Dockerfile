FROM node:16-alpine AS frontend
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM golang:1.19 AS builder
WORKDIR /go/app
COPY . .
RUN go build
