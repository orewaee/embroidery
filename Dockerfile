FROM golang:alpine AS BUILD
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /embroidery ./cmd/embroidery/main.go

FROM alpine AS RELEASE

COPY --from=BUILD /embroidery /embroidery

EXPOSE 8080

ENTRYPOINT ["/embroidery"]

