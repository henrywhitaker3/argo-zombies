FROM golang:1.22 as builder

WORKDIR /builder

COPY . /builder/

RUN go mod download

RUN CGO_ENABLED=0 go build -a -o argo-zombies main.go

FROM scratch

COPY --from=builder /builder/argo-zombies /

ENTRYPOINT [ "/argo-zombies" ]
