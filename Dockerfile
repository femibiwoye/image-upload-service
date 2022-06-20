# Start from golang base image
FROM golang:1.18-alpine as builder

WORKDIR /app

COPY . .

ENV GO111MODULE=on

RUN go build -o server -ldflags='-s -w' ./cmd/image-uploader/main.go

# Add Maintainer info
FROM golang:1.18-alpine

# creates working directory for program
WORKDIR /app

COPY --from=builder /app/server .

COPY assets  ./assets

CMD [ "./server" ]