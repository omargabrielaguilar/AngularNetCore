# Base image
FROM golang:1.21

# Set working directory
WORKDIR /app

# Copy files
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

# Expose the port
EXPOSE 8080

# Run the app
CMD ["./main"]
