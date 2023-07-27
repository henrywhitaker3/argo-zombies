FROM scratch
WORKDIR /
COPY argo-zombies /argo-zombies
USER 65532:65532

ENTRYPOINT ["/argo-zombies"]
