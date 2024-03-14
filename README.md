# REST API with Go and Docker

This project is a simple trivia REST API built using Go and Docker. The API is built using the Go Fiber framework and connected to a Postgres database.

## Prerequisites

To run this project, you will need to have Docker installed on your machine.

## Getting Started

1. Clone this repository to your local machine.
2. Change into the project directory: `cd --name of the folder--`
3. Build the Docker images: `docker compose build`
4. Start the containers: `docker compose up`
5. The API will be available at `http://localhost:3000`.

## API Endpoints

The following endpoints are available:

* `GET /`: Returns a welcome message.
* `GET /facts`: Returns a list of all facts.
* `POST /fact`: Creates a new fact. The request body should be a JSON object with `question` and `answer` fields.

## Project Structure

The project is structured as follows:

* `cmd/main.go`: The entry point for the application.
* `models/models.go`: Defines the `Fact` struct and its associated database schema.
* `database/database.go`: Handles the connection to the Postgres database.
* `handlers/facts.go`: Defines the API endpoints and their corresponding handler functions.
* `Dockerfile`: The Dockerfile used to build the Go application.
* `docker-compose.yml`: The Docker Compose file used to define the application's services and their dependencies.

## Environment Variables

The following environment variables are used in this project:

* `DB_USER`: The username for the Postgres database.
* `DB_PASSWORD`: The password for the Postgres database.
* `DB_NAME`: The name of the Postgres database.

These variables should be defined in a `.env` file in the root of the project directory.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.