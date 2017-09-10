# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/andrewcopp/andrewcopp/resource/posts

# Build the posts command inside the container.
RUN go get github.com/lib/pq
RUN go install github.com/andrewcopp/andrewcopp/resource/posts

# Run the posts command by default when the container starts.
ENTRYPOINT /go/bin/posts

# Document that the service listens on port 8080.
EXPOSE 8080
