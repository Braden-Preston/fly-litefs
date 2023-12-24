FROM golang:1.21

# Copy source code (w/.dockerignore)
COPY ./ /app

# Install system dependencies
RUN apt-get update -y && \
    apt-get install -y ca-certificates fuse3 sqlite3 curl

# Set working directory
WORKDIR /app

# Build golang packages and utitlies
RUN go mod download
RUN go build -o ./bin/app
RUN chmod -R ug+x ./bin

# Add working directory binaries to path
ENV PATH="${PATH}:/app/bin"

# Configure application
ENV LOG_LEVEL = "DEBUG"

# Open up app ports to host
EXPOSE 4000

# Install LiteFS
COPY --from=flyio/litefs:0.5 /usr/local/bin/litefs /usr/local/bin/litefs

# Use LifeFS as supervisor and entrypoint
ENTRYPOINT [ "litefs", "mount" ]
