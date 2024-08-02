# Build the Docker image
docker build -t analytic-service .

# Run the Docker container
docker run -p 8080:8080 analytic-service
