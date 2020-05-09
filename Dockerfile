# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.13

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/ahsu1230/GolangDockerTest
WORKDIR /go/src/github.com/ahsu1230/GolangDockerTest

# Install project dependencies
RUN go get github.com/ahsu1230/GolangDockerTest

# Build the command inside the container by installing the binary.
RUN go install .

# Run the command by default when the container starts.
ENTRYPOINT /go/bin/GolangDockerTest

# Document that the container service listens on port 8080.
EXPOSE 8080