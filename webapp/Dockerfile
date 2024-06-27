# Step 1: Build the Go binary in a builder stage
FROM golang:1.22 as builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app as a static binary
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o horner cmd/serve/main.go

# Step 2: Use scratch as final base image, copy the static binary from builder
FROM scratch

EXPOSE 8888
# Copy the binary to scratch
COPY --from=builder /app/horner /horner

# Command to run the executable
CMD ["/horner"]
