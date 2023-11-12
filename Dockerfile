FROM golang:1.20 as builder

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod .

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Set the entrypoint command to run the script
ENTRYPOINT ["./main"]