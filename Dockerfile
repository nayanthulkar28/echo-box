FROM golang:1.23.3-alpine3.19
WORKDIR /workspace
COPY cmd cmd
COPY config config
COPY internal internal
COPY pkg pkg
COPY go.mod go.mod
COPY go.sum go.sum
COPY .env.production .env
EXPOSE 8080
RUN go build cmd/main.go
ENTRYPOINT [ "./main" ]