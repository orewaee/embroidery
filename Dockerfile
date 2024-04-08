FROM golang:alpine AS BUILD

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o embroidery-api main.go

FROM alpine

WORKDIR /build

COPY --from=BUILD /build/embroidery-api /build/embroidery-api
COPY .env .
COPY files .

RUN mkdir logs
CMD ["./embroidery-api"]
