# Use the official Golang image as a base
FROM golang:1.23.3

# Set the current working directory inside the container
WORKDIR /usr/local/src

# Copy the source code into the container
COPY . .

# Compile code
RUN go build -o /usr/local/bin/gateway/gateway /usr/local/src/cmd/main.go

# Ensure the binary has executable permissions
RUN chmod +x /usr/local/bin/gateway/gateway

# Set the default command to run the compiled binary
CMD ["/usr/local/bin/gateway/gateway"]