# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.22.2-alpine

LABEL org.opencontainers.image.authors="Josh Ellithorpe <quest@mac.com>"

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/zquestz/s

# Switch to the correct working directory.
WORKDIR /go/src/github.com/zquestz/s

# Build the code.
RUN apk add --update make
RUN make install

# Set the start command.
ENTRYPOINT ["/go/bin/s", "-s"]

# Document that the service listens on port 8080.
EXPOSE 8080
