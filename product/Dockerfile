FROM golang:1.13.4-alpine3.10 AS build
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/product
COPY . .
RUN go mod download
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/product ./main.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/product/bin /go/bin
WORKDIR /go/bin
ADD .env /go/bin
RUN mkdir files
EXPOSE 9091
ENTRYPOINT /go/bin/product --port 9091