# URL Shortener Application

This is a Golang-based URL shortening service that provides:
- REST APIs to create and manage shortened URLs
- A web interface to interact with the service
- Swagger documentation for API exploration

## Getting Started

### 1. Install Golang

Ensure you have Go installed (version 1.20+ recommended):

```bash
brew install go
```

Verify installation:

```bash
go version
```

### 2. Run the Project

Navigate to the project directory and start the application:

```bash
cd /Users/anuragupadhyay/Desktop/Anurag/projects/shortenurl
make run
```

This will:
- Install dependencies
- Build the application
- Start the server on port 8080

### 3. Explore APIs

Access the Swagger documentation at: http://localhost:8080/swagger/index.html
Here you can:
- View all available API endpoints
- Test API calls directly
- See request/response formats

### 4. Use the Web Interface

Open the application in your browser: http://localhost:8080
The web interface allows you to:
- Create shortened URLs

The application runs with a local SQLite database by default, requiring no additional setup. For production use, you may want to configure a different database in the application configuration.

