# Autocomplete API

A Go-based REST API that provides country name autocomplete functionality. The API uses a Trie data structure for efficient prefix-based searching and matching.

## Features

- Fast autocomplete suggestions for country names
- Case-insensitive searching
- RESTful API endpoints
- Built with Echo framework
- Efficient Trie-based search implementation

## API Endpoints

- `GET /` - Welcome message
- `GET /countries?q={query}` - Get country suggestions based on the query parameter

## Requirements

- Go 1.24.3 or higher

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/autocomplete.git
cd autocomplete
```

2. Install dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run server.go
```

The server will start on port 1323.

## Usage

To get country suggestions, send a GET request to `/countries` with a query parameter:

```
http://localhost:1323/countries?q=aus
```

This will return matches like "Australia", "Austria", etc.

## Project Structure

- `server.go` - Main application entry point
- `controllers/` - Request handlers and business logic
- `pkg/trie/` - Trie data structure implementation

## License

This project is licensed under the MIT License - see the LICENSE file for details.
