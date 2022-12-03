###############################################################################
# Modules caching
###############################################################################
FROM golang:1.19.3-alpine3.17 as modules
LABEL maintainer="jtbonhomme@gmail.com"
# We want to populate the module cache based on the go.{mod,sum} files. 
COPY go.mod go.sum /modules/
WORKDIR /modules
# Because of how the layer caching system works in Docker, the go mod download 
# command will _ only_ be re-run when the go.mod or go.sum file change 
# (or when we add another docker instruction this line) 
RUN go mod download


###############################################################################
# Builder
###############################################################################
FROM golang:1.19.3 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o /bin/app \
  -ldflags "-X github.com/jtbonhomme/go-template/internal/version.Tag=$(git describe --tags --match '[0-9]*\.[0-9]*\.[0-9]*' 2> /dev/null || echo 'no-tag') \
  -X github.com/jtbonhomme/go-template/internal/version.GitCommit=$(git rev-parse --short HEAD) \
  -X github.com/jtbonhomme/go-template/internal/version.BuildTime=$(date -u +%FT%T%z)" \
  ./cmd/app

###############################################################################
# Run program and expose server listening port
###############################################################################
FROM golang:1.19.3-alpine3.17

# This environment variable will overwrite default configuration
ENV PORT=8090
EXPOSE 8090

COPY --from=builder /app/config.yml /config.yml
COPY --from=builder /bin/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /
CMD ["/app"]
