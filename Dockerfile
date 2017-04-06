# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

MAINTAINER Josh Ellithorpe <quest@mac.com>

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/zquestz/s

# Switch to the correct working directory.
WORKDIR /go/src/github.com/zquestz/s

# Restore vendored packages.
RUN go get -u github.com/FiloSottile/gvt
RUN gvt restore

# Build the code.
RUN make install

# Set the start command.
ENTRYPOINT ["/go/bin/s", "-s"]

# Document that the service listens on port 8080.
EXPOSE 8080
