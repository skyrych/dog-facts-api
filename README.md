# Dog Facts API

[](https://golang.org/)
[](https://opensource.org/licenses/Apache-2.0)
[](https://www.docker.com/)

A simple, secure, and containerized web service written in Go that serves fascinating facts about a dog named Needy.

This project serves as a practical example of building a Go application with modern best practices, including structured logging, HTTPS, and secure multi-stage Docker builds using distroless images.

-----

## Features

  * **RESTful Endpoint**: Provides a single `/facts` endpoint to retrieve a random dog fact.
  * **Secure by Default**: Communicates exclusively over HTTPS.
  * **Structured Logging**: Uses Go's standard `log/slog` library for structured JSON logs, which is ideal for production environments.
  * **Containerized**: Includes a multi-stage `Dockerfile` that builds a minimal, secure, non-root container using distroless images.
  * **Graceful Shutdown**: The server is configured to shut down gracefully, ensuring in-flight requests are completed.

-----

## Prerequisites

  * [Go](https://go.dev/doc/install) (version 1.21 or higher)
  * [Docker](https://www.docker.com/products/docker-desktop)

-----

## Getting Started

### 1\. Clone the Repository

```bash
git clone https://github.com/skyrych/dog-facts-api.git
cd dog-facts-api
```

### 2\. Build and Run the API (Locally)

This runs the application directly from source.

```bash
# Optional: Set the PORT environment variable. Defaults to :8080 if not set.
export PORT=:8080

# Run the application from the main package directory
go run ./cmd/dogfacts
```

### 3\. Build and Run the API (in a Docker Container)

This runs the application inside its Docker container.

```bash
# Build the Docker image
docker build -t dog-facts-api:v1.0.0 .

# Run the container, mapping port 8080 on the host to port 80 in the container
docker run -p 8080:80 dog-facts-api:v1.0.0
```

-----

## API Endpoints

The API is available on `http://localhost:8080` (or your configured port) and provides the following endpoints:

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/facts` | Returns a random fact about Needy the dog. |
| `GET` | `/healthz` | A simple health check endpoint. Returns `200 OK` if the server is running. |

-----

## Testing

To run the unit tests for the project's core logic:

```bash
go test ./internal/app/dogfacts -v
```

-----

## Technologies Used

  * **Go**: The programming language for the API server.
  * **`net/http`**: Go's built-in package for HTTP networking.
  * **`log/slog`**: Go's standard library for structured logging.
  * **`testing`**: Go's built-in framework for unit testing.
  * **Docker**: For containerizing the application.