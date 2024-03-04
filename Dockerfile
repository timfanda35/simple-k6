# Builder
FROM golang:1.21 as builder

WORKDIR /src
COPY . ./
RUN make build

# Builder
FROM grafana/k6:0.49.0

WORKDIR /

EXPOSE 8080
EXPOSE 5665

COPY --from=builder /src/app /

ENTRYPOINT ["/app"]
