# Go Echo Demo API

This repository contains a simple Go implementation of an API that uses the Echo web framework to handle HTTP requests. The API has endpoints for fetching posts, as well as a health check endpoint. The implementation includes handlers for each of these endpoints, which interact with a service layer that provides the necessary functionality.

## Installation

Clone the repository and run `go mod download` to download the necessary dependencies.

## Usage

Run the application using the command `go run main.go`. The application will start running on port 1323.

## Endpoints

### GET /health-check

This endpoint returns a simple message indicating whether the API is up and running.

### GET /posts

This endpoint returns a list of posts in JSON format.

### GET /post/:id

This endpoint returns a single post based on the ID specified in the URL.

## Testing

The repository includes a test suite that tests the endpoints of the API. Run the tests using the command `go test ./...`. The test suite includes tests for the handlers that fetch posts as well as a health check endpoint.
