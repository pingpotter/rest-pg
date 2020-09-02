# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.15.1-alpine3.12 as stage1

# Add Maintainer Info
LABEL maintainer="jiraphon sa."

# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY cmd cmd
COPY internal internal

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o rest-pg ./cmd

# generate clean, final image for end users
FROM alpine:3.12
COPY --from=stage1 /app/rest-pg .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./rest-pg"]