# ! First stage

FROM golang:latest AS build

# Set the working directory inside the container
WORKDIR /app

# Download and install any required dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/soccerfriend .


# ! Final stage

FROM alpine:latest

WORKDIR /app

# Copy only the necessary files from the build stage
COPY --from=build /app/bin/soccerfriend .
#COPY --from=build /app/.env .

ENV TOKEN=""
ENV APIKEY=""

# Command to run the app
CMD ["/app/soccerfriend"]
