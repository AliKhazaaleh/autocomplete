# Autocomplete API

A Go-based REST API that provides intelligent autocomplete functionality for various categories including countries, cities, universities, companies, job titles, and skills. The API uses a Trie data structure for efficient prefix-based searching and matching.

## Features

- Fast autocomplete suggestions for multiple categories
- Case-insensitive searching
- RESTful API endpoints
- Built with Echo framework
- Efficient Trie-based search implementation

## API Endpoints

- `GET /` - Welcome message
- `GET /countries?q={query}` - Get country name suggestions
- `GET /cities?q={query}` - Get city name suggestions
- `GET /universities?q={query}` - Get university name suggestions
- `GET /companies?q={query}` - Get company name suggestions
- `GET /job-titles?q={query}` - Get job title suggestions
- `GET /skills?q={query}` - Get skill suggestions

All endpoints support the following:
- Case-insensitive searching
- Partial matching
- Returns all items when query parameter is empty
- Returns matching items sorted by relevance when query is provided

Example usage:
# Get countries containing "United"
GET http://localhost:1323/countries?q=United

# Get cities starting with "San"
GET http://localhost:1323/cities?q=San

# Get universities containing "University of"
GET http://localhost:1323/universities?q=University of

# Get companies containing "Inc"
GET http://localhost:1323/companies?q=Inc

# Get job titles containing "Engineer"
GET http://localhost:1323/job-titles?q=Engineer

# Get skills containing "Python"
GET http://localhost:1323/skills?q=Python
```

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
