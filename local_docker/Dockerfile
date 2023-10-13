# Use the official Go image as a parent image
FROM golang:1.17

# Set the current working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY ./http_server.go ./http_server.go

# Build the Go app
RUN go build -o server http_server.go

# Expose port 8081
EXPOSE 8081

# Command to run the binary
CMD ["./server"]
