# Start from the official GoLang image, specifying the version you want
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download and install Go dependencies
RUN go mod download

# Copy the rest of the project files
COPY . .

# Build the Go application
RUN go build -o main .

# Set the entry point for the container
ENTRYPOINT ["./main"]
