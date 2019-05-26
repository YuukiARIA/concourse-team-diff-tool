FROM golang:1.12 AS builder

COPY . /workspace
WORKDIR /workspace

RUN CGO_ENABLED=0 go build -o bin/glanceable

FROM busybox

COPY --from=builder /workspace/bin/glanceable /usr/local/bin/glanceable

ENTRYPOINT ["/usr/local/bin/glanceable"]
