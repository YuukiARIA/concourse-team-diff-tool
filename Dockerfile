FROM golang:1.12 AS builder

COPY . /workspace
WORKDIR /workspace

RUN CGO_ENABLED=0 go build -o bin/team-diff

FROM busybox

COPY --from=builder /workspace/bin/team-diff /usr/local/bin/team-diff

ENTRYPOINT ["/usr/local/bin/team-diff"]
