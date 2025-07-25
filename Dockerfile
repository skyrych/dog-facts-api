# Stage 1: Build the Go application
# Use a Go base image with a specific version (e.g., 1.22-alpine for smaller images, or latest if preferred)
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files (go.mod and go.sum)
# This allows Docker to cache the 'go mod download' step if your dependencies don't change
COPY go.mod ./

# Download all Go module dependencies
# This pulls in necessary packages for your application
RUN go mod download

# Copy the entire source code into the container
# The '.' refers to the current directory on your host (the project root)
COPY . .

# Build the Go application binary
# CGO_ENABLED=0: Disables CGO, producing a statically linked binary (important for distroless/scratch)
# GOOS=linux: Compiles for Linux operating system (standard for containers)
# GOARCH=amd64: Compiles for AMD64 architecture
# -o /app/dog-facts-api: Specifies the output path and name for the compiled binary
# ./cmd/dogfacts: Points to the directory containing your main package
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/dog-facts-api ./cmd/dogfacts

# Stage 2: Create the final, minimal image
# Use a distroless image for maximum security and minimal size.
# 'nonroot' variant runs as a non-root user by default.
FROM gcr.io/distroless/static:nonroot

# Set the working directory in the final image
WORKDIR /

# Copy the compiled binary from the 'builder' stage into the root of the final image
# /app/dog-facts-api is the path inside the 'builder' stage where the binary was created.
# /dog-facts-api is the path in the final image where the binary will reside.
COPY --from=builder /app/dog-facts-api /dog-facts-api

# Expose the port your application listens on.
# This is metadata; you still need to map the port when running the container.
EXPOSE 8080

# Define the command to run when the container starts.
# This should point to your compiled binary.
ENTRYPOINT ["/dog-facts-api"]