FROM golang:1.13-stretch as builder
WORKDIR /go/src/code.cloudfoundry.org/eirini-staging
COPY . .
RUN GO111MODULE=on GOOS=linux go build -mod vendor -a -o /executor cmd/executor/executor.go

FROM cloudfoundry/cflinuxfs3

WORKDIR /workspace

RUN mkdir -p /packs

COPY --from=builder /executor /packs/

ENTRYPOINT [ \
  "/packs/executor" \
]
