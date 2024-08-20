FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod download

RUN go install github.com/CycloneDX/cyclonedx-gomod/cmd/cyclonedx-gomod@latest

# Keep the container running and allow shell access
CMD ["tail", "-f", "/dev/null"]