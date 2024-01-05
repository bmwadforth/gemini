FROM golang:1.21.3

ENV PORT=8080
ENV APP_ENV=PRODUCTION
# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
#COPY *.go ./
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /svc-template

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/svc-template"]