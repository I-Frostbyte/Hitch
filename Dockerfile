FROM golang:1.25-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy only the go.mod and go.sum files to leverage Docker cache for dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install curl
RUN apk add --no-cache curl

# Copy the code into the container
COPY . .

# Build binary into a dedicated Locations (IMPORTANT FIX)
RUN mkdir -p /app/bin
RUN go build -o /app/bin/server ./users/cmd/grpc-service

# Expose the gRPC server port
EXPOSE 40041

CMD ["/app/bin/server"]