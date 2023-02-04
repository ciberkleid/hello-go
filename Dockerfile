FROM golang
WORKDIR /workspace
COPY go.mod .
RUN go mod download -json
COPY hello-server.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -installsuffix cgo -o hello-server

FROM scratch
COPY --from=0 /workspace/hello-server /
USER 1001
ENTRYPOINT ["/hello-server"]
