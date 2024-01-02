# Stage 1: Build stage
FROM golang:1.21-alpine3.19 as builder

# Set necessary environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Install dependencies
RUN apk add --update curl

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy the Go module and Go sum files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the application
RUN go build ./services/contact/cmd/app

# Stage 2: Final stage
FROM scratch

# Set the working directory inside the container
WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /usr/src/app/app .

# Command to run the application
CMD [ "/app" ]
