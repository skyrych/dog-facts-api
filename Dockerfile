# Stage 1: Build the Go application
# Use a Go base image with a specific version (e.g., 1.22-alpine for smaller images, or latest if preferred)
FROM golang:1.24.3 AS builder

# Set the working directory inside the container for the build process
WORKDIR /app

# Copy the Go module files (go.mod and go.sum)
# This allows Docker to cache the 'go mod download' step if your dependencies don't change
COPY go.mod ./

# Download all Go module dependencies
RUN go mod download

# Copy the entire source code into the container
# The '.' refers to the current directory on your host (the project root)
COPY . .

# Build the Go application binary
# CGO_ENABLED=0: Disables CGO, producing a statically linked binary
# GOOS=linux GOARCH=amd64: Compiles for Linux AMD64 (standard for containers)
# -o /app/dog-facts-api: Specifies the output path and name for the compiled binary within this stage
# ./cmd/dogfacts: Points to the directory containing your main package
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/dog-facts-api ./cmd/dogfacts

# Stage 2: Create the final, minimal, secure image
# Use a distroless image for maximum security and minimal size (no shell, no package manager)
# 'static:nonroot' is suitable for statically linked Go binaries and runs as non-root user.
FROM gcr.io/distroless/static:nonroot

# Set the working directory in the final image to /app
# This means your binary will be at /app/dog-facts-api in the final image
WORKDIR /app

# Copy the compiled binary from the 'builder' stage into the final image
# Syntax: COPY --from=<stage_name> <source_path_in_stage> <destination_path_in_final_image>
COPY --from=builder /app/dog-facts-api /app/dog-facts-api


# Expose the standard HTTP port (80)
# This is metadata informing Docker that the container expects traffic on this port.
EXPOSE 80

# Define the command to run when the container starts.
# This should point to your compiled binary within the final image.
ENTRYPOINT ["/app/dog-facts-api"]
