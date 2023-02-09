# Use the latest golang image as the base
FROM golang:latest

# Set the working directory in the container to /app
WORKDIR /app

# Copy the Go source code to the container /app directory
COPY . /app

# Compile the Go code into an executable
RUN go build main.go

EXPOSE 8000

# Set the command to run when the container starts
CMD ["./main"]

