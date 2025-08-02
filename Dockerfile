FROM golang:1.18

WORKDIR /app

# Copy go.mod first to leverage Docker cache
COPY go.mod ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the server application
RUN go build -o app .

EXPOSE 80

# Run the server application
CMD [ "./app" ]
