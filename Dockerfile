# Use the official Golang image as a base
FROM golang:1.23.3

# Set the current working directory inside the container
WORKDIR /usr/local/src

# Copy the source code into the container
COPY . .

# Compile code
# RUN go build -o /usr/local/bin/gateway/gateway main.go

# Set the default command to run the compiled binary
# CMD ["/usr/local/bin/gateway/gateway"]

CMD ["tail", "-f", "/dev/null"]