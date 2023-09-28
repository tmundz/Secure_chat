#Base  go-image
FROM golang:1.21.1-bookworm

# Create a directory for the app
RUN mkdir /app

# Copy all file from current directory
COPY . /app

# Set working directory
WORKDIR /app

# build the go executable
RUN go build -o server .


# Run the server executable
CMD ["/app/server"]