FROM golang:1.21-alpine as certs

RUN apk --no-cache add ca-certificates

FROM scratch
WORKDIR /

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY argo-zombies /argo-zombies
USER 65532:65532

ENTRYPOINT ["/argo-zombies"]
