# Start from golang v1.11 base image
FROM golang:1.11

# Add Maintainer Info
LABEL maintainer="Hilman Syafei <hilman@alterra.id>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/hilmansyafei/canopus-master-mservice

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Install dep
RUN go get -u github.com/golang/dep/cmd/dep

# install 
RUN dep ensure

# This container exposes port 8080 to the outside world
EXPOSE 1324

# Run the executable
CMD go run main.go