# Use Golang Version
FROM golang:latest

# Working Directory Setting
WORKDIR /server

# Base File Copy
COPY go.mod go.sum ./
RUN go mod download

# Container의 Working Directory로 Code Copy
COPY . .

# Build
RUN go build -o main .

# Execute
CMD ["./main"]