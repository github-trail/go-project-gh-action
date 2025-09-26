# Go REST API

This project is a simple REST API built using Go. It demonstrates how to handle HTTP GET requests and serves a JSON response.

## Project Structure

```
go-test-api
├── src
│   ├── main.go          # Entry point of the application
│   ├── handlers
│   │   └── get_handler.go # Handles HTTP GET requests
│   └── routes
│       └── routes.go    # Defines the routes for the application
├── .github
│   └── workflows
│       └── docker-push.yml # GitHub Actions workflow for Docker
├── Dockerfile            # Dockerfile for building the Docker image
├── go.mod                # Go module definition
└── README.md             # Project documentation
```

## Getting Started

### Prerequisites

- Go 1.24.3 or later
- Docker
- Docker Hub account (for pushing the Docker image)

### Running the Application

1. Clone the repository:
   ```
   git clone <repository-url>
   cd go-test-api
   ```

2. Build the application:
   ```
   go build -o main src/main.go
   ```

3. Run the application:
   ```
   ./main
   ```

4. Access the API:
   Open your browser or use a tool like `curl` to access the API at `http://localhost:8080`.

### Docker

To build and run the Docker image:

1. Build the Docker image:
   ```
   docker build -t <your-dockerhub-username>/go-test-api .
   ```

2. Run the Docker container:
   ```
   docker run -p 8080:8080 <your-dockerhub-username>/go-test-api
   ```

### GitHub Actions

This project includes a GitHub Actions workflow that automatically builds and pushes the Docker image to Docker Hub whenever changes are pushed to the main branch. The workflow is defined in `.github/workflows/docker-push.yml`.

## API Endpoints

- `GET /`: Returns a simple JSON message.

## License

This project is licensed under the MIT License.