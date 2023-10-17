FROM golang:1.20 AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

FROM alpine:latest
WORKDIR /app


COPY --from=build /app .



# Copy the built Go binary from the build stage to the runtime stage

# Expose port 8080 to the outside world
EXPOSE 8080

# Set the entry point for the container
CMD ["/app/main"]