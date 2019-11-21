FROM golang:latest as builder

ADD . /go/src/github.com/ViniciusAyres/user-API
WORKDIR /go/src/github.com/ViniciusAyres/user-API
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /build/app .

FROM scratch
COPY --from=builder /build/app/ .
ENTRYPOINT ["/app"]