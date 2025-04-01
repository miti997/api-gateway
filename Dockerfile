# Use the official Golang image as a base
FROM golang:1.23.3

# Set the current working directory inside the container
WORKDIR /usr/local/src

# Copy the source code into the container
COPY . .

# Copy the config folder from /usr/local/ to /usr/local/bin/gateway/
COPY config /usr/local/bin/gateway/config

# Remove the .example extension from all files in the config folder
RUN find /usr/local/bin/gateway/config -type f -name "*.example" -exec bash -c 'mv "$0" "${0%.example}"' {} \;

# Compile code
RUN go build -o /usr/local/bin/gateway/gateway /usr/local/src/cmd/main.go

# Ensure the binary has executable permissions
RUN chmod +x /usr/local/bin/gateway/gateway

#Keep container runing to have time to change configs
CMD ["tail", "-f", "/dev/null"]
