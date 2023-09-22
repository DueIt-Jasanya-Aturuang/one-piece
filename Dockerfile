# Stage 1: Build the Go application
FROM golang:1.21 AS build
# Set cgo_enabled to 1
ENV CGO_ENABLED=0

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the Go application
RUN go build -o finance .

# Stage 2: Create a minimal runtime container using Alpine Linux
FROM alpine:latest AS final

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=build /app/finance .
#
## Install GLIBC
#RUN apk add gcompat

# Expose the port your application listens on (if applicable)
# EXPOSE 8080

# Run your application
CMD ["./finance"]