FROM golang:alpine AS build-env

# Set working directory for the build
WORKDIR /go/src/github.com/marbar3778/do_or_dare

# Setup build environment
RUN apk add --no-cache curl git && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Copy only the dependency manifests
COPY Gopkg.toml Gopkg.lock ./

# Fetch dependencies
RUN dep ensure -v --vendor-only

# Build dependencies
RUN find vendor -maxdepth 2 -mindepth 2 -type d -exec \
    sh -c 'CGO_ENABLED=0 go build -v -ldflags "-s -w" github.com/marbar3778/do_or_dare/{}/... || true' \;

# Add source files
COPY . ./

# Build and install
RUN \
    CGO_ENABLED=0 go build -v -ldflags "-s -w" -o build/do_or_dared ./cmd/do_or_dared && \
    CGO_ENABLED=0 go build -v -ldflags "-s -w" -o build/do_or_darecli ./cmd/do_or_darecli

# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/src/github.com/marbar3778/do_or_dare/build/do_or_dared /usr/bin/do_or_dared
COPY --from=build-env /go/src/github.com/marbar3778/do_or_dare/build/do_or_darecli /usr/bin/do_or_darecli

# Run the daemon by default
CMD ["do_or_dared"]
