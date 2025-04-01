# API Gateway

This project provides a simple API gateway using Docker for containerized deployment and management.

## Getting Started

### Prerequisites
- Make sure you have [Docker](https://www.docker.com/) installed.
- Install [Make](https://www.gnu.org/software/make/).

### Running the API Gateway

Use the following `make` commands to manage the gateway:

- **Start the container**:
  ```sh
  make up
  ```
  This command will build and start the API gateway container.

- **Recreate the container**:
  ```sh
  make recreate
  ```
  This stops, removes, and rebuilds the container from scratch.

- **Start the server inside the container**:
  ```sh
  make start
  ```
  This will start the API server if the container is already running.

- **Enter the container's bash shell**:
  ```sh
  make enter
  ```
  This command opens a shell session inside the running container.

## License
This project is licensed under the MIT License.

