FROM golang:latest as builder
# ARG security: https://bit.ly/2oY3pCn
ARG SSH_PRIVATE_KEY
RUN mkdir -p /root/.ssh && umask 0077 echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa 
RUN git config --system url."ssh://git@github.com/".insteadOf "https://github.com/"
RUN ssh-keyscan -H github.com >> /root/.ssh/known_hosts
ADD . /go/src/github.com/ViniciusAyres/user-API
WORKDIR /go/src/github.com/ViniciusAyres/user-API
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /build/app .

FROM scratch
COPY --from=builder /build/app/ .
ENTRYPOINT ["/app"]