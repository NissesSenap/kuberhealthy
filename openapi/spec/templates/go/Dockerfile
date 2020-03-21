FROM golang AS builder
ADD . /build
WORKDIR /build/my-external-check
RUN CGO_ENABLED=0 go build -v

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/my-external-check /app/my-external-check 
ENTRYPOINT ["/app/my-external-check "]
