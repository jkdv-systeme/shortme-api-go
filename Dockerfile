# build stage
FROM golang:1.18-alpine AS build

RUN apk --no-cache add ca-certificates

WORKDIR /build

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o short-me .


# run stage
FROM scratch

WORKDIR /opt

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /build/short-me .

ENTRYPOINT ["./short-me", "serve"]