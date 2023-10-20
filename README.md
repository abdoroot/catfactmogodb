# CatFact API and Database

A simple Go project that fetches cat facts from an external API, stores them in a local MongoDB database, and provides an HTTP endpoint to retrieve these cat facts.

## Usage

1. Start your MongoDB server on `mongodb://localhost:27017`.

2. Run the Go application:

   ```bash
   go run main.go
   ```

3. Retrieve cat facts via a GET request to `http://localhost:3000/facts`.

## Endpoints

- `GET /facts`: Retrieve stored cat facts in JSON format.

## Configuration

Modify MongoDB connection URI and HTTP server port in `main.go` if needed.

## Contributing

Feel free to contribute or report issues on GitHub. Enjoy your cat facts! 🐱🐾