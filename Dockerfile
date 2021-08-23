FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GIN_MODE=release \
    DATABASE_USER=admin \
    DATABASE_PASS=wXSVJo2I \
    DATABASE_PORT=18809 \
    DATABASE_HOST=mysql-45996-0.cloudclusters.net \
    DATABASE_NAME=productapi \
    JWT_SECRET=ih3i1uh31wqc 

# Create app directory
WORKDIR /app

# Copy all other source code to work directory
COPY . .

# Download all the dependencies that are required
RUN go mod tidy

# Build the application
RUN go build -o binary main.go

ENTRYPOINT ["/app/binary"]