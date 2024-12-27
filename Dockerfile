# Specifies a parent image
FROM golang:1.23

# Create a directory for the app that holds the app code
WORKDIR /app

# Copy the app code into the container
COPY . .

# Download all dependencies
RUN go mod download

# Build the app
RUN go build -o /go-docker

# Expose the port the app runs on
EXPOSE 8181

# Run the app
CMD [ "/go-docker" ]

