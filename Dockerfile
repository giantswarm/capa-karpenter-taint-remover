# The binary is pre-built by architect/go-build in CircleCI (matrix over
# linux/amd64 and linux/arm64) and persisted to the workspace as
# capa-karpenter-taint-remover-linux-<arch>. docker buildx sets TARGETARCH
# per platform; we copy the matching one into the image.
FROM gcr.io/distroless/static:nonroot
WORKDIR /
ARG TARGETARCH
COPY capa-karpenter-taint-remover-linux-${TARGETARCH} /capa-karpenter-taint-remover
USER 65532:65532

ENTRYPOINT ["/capa-karpenter-taint-remover"]
