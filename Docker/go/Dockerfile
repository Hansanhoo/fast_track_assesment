# syntax=docker/dockerfile:1

FROM golang:1.21

# Set working directory inside the container
WORKDIR /app

RUN apt-get update && apt-get install -y curl \
    && curl -o /wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh \
    && chmod +x /wait-for-it.sh

# Copy go mod files and vendor directory
COPY go.mod go.sum ./
COPY vendor/ ./vendor/
# Enable vendoring mode
ENV GOFLAGS="-mod=vendor"

# Copy the rest of the source code
COPY . .

# Build the Go app from standard layout
RUN go build -o /main ./cmd/assesment_fast_track

# Run the binary
CMD ["/wait-for-it.sh", "rabbitmq:5672", "--", "/wait-for-it.sh", "mysql:3306", "--", "/main"]
