# API Gateway

This project provides a simple API Gateway using Docker for containerized deployment and management. It acts as a proxy, routing API requests to specified outbound URLs while preserving headers, parameters, and responses.

## Getting Started

### Prerequisites

Ensure that you have the following installed:

- [Docker](https://www.docker.com/)
- [Make](https://www.gnu.org/software/make/)

### Running the API Gateway

Use the following `make` commands to manage the gateway:

- **Start the container**:
  ```sh
  make up
  ```
  This command builds and starts the API Gateway container.

- **Recreate the container**:
  ```sh
  make recreate
  ```
  Stops, removes, and rebuilds the container from scratch.

- **Start the server inside the container**:
  ```sh
  make start
  ```
  Starts the API server if the container is already running.

- **Enter the container’s bash shell**:
  ```sh
  make enter
  ```
  Opens a shell session inside the running container.

- **Run tests**:
  ```sh
  make test
  ```

- **Run tests in verbose mode**:
  ```sh
  make test-v
  ```

### Configuration Files

The project comes with example configuration files. These files are copied to `/usr/local/bin/gateway/config` when the container starts.

#### `config.json`

This is the main configuration file for the API Gateway. It defines the server's basic configuration, such as the address and port.

Example:
```json
{
    "addr": ":8080"
}
```

#### `routing.json`

This file holds the route definitions for the API Gateway. It maps incoming requests to outbound URLs, forwarding headers, parameters, and response data.

Example:
```json
[
    {
        "request": "GET",
        "in": "/posts",
        "out": "https://jsonplaceholder.typicode.com/posts"
    },
    {
        "request": "GET",
        "in": "/posts/{id}",
        "out": "https://jsonplaceholder.typicode.com/posts/{id}"
    },
    {
        "request": "POST",
        "in": "/posts",
        "out": "https://jsonplaceholder.typicode.com/posts"
    }
]
```

#### `logger_config.json`

This file configures the logging behavior of the API Gateway. It specifies log file settings, including maximum file size, file path, and file name.

Example:
```json
{
    "maxSizeMB": 1024,
    "filePath": "/usr/local/bin/gateway/logs",
    "fileName": "log"
}
```

### License

This project is licensed under the MIT License.