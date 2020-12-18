FROM golang AS builder

# enable Go modules support
ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/tekton/pipeline

# manage dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy src code from the host and compile it
COPY cmd cmd
COPY pkg pkg
COPY vendor vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /webhook cmd/webhook/main.go

FROM alpine:3.9
RUN apk --no-cache add ca-certificates
COPY --from=builder /webhook /bin

# Expose port 8080 to the outside world
EXPOSE 8080

ENTRYPOINT ["/bin/webhook"]
