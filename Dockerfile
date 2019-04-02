# Build the sockmap binary
FROM golang:1.12.1 as builder

# Copy in the go src
WORKDIR /go/src/github.com/dippynark/sockmap
COPY pkg/    pkg/
COPY cmd/    cmd/
COPY vendor/ vendor/

# Build
# TODO: statically compile CGO_ENABLED=0 
RUN GOOS=linux GOARCH=amd64 go build -tags netgo -a -o sockmap github.com/dippynark/sockmap/cmd/sockmap

# Copy the sockmap binary into a thin image
FROM frolvlad/alpine-glibc
WORKDIR /
COPY --from=builder /go/src/github.com/dippynark/sockmap/sockmap .
ENTRYPOINT ["/sockmap"]
