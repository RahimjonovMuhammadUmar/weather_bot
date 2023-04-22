# Build stage

FROM golang:latest AS builder

# Set the working directory
WORKDIR /app

# Copy all the application code into the container
COPY . .

RUN go mod tidy

# Install the necessary dependencies for database migration
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Compile the application and save the executable file inside the container
RUN go build -o main cmd/main.go


# Run stage

FROM alpine:latest

# Set the working directory for running 
WORKDIR /app

# Copy the executable file `main` from the build stage to the run stage
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["/app/main"]
