# Autocomplete API

A Go-based REST API that provides intelligent autocomplete functionality for various categories including countries, cities, universities, companies, job titles, and skills. The API uses a Trie data structure for efficient search.

## Features

- Fast autocomplete suggestions for multiple categories
- Case-insensitive searching
- RESTful API endpoints
- Built with Echo framework
- Efficient Trie-based search implementation
- Configurable result limits

## API Endpoints

All endpoints support the following query parameters:
- `q`: Search query string
  - Case-insensitive matching
  - Partial string matching
- `limit`: Maximum number of results (optional)
  - Default: 20
  - Maximum: 20
  - Must be positive integer

### Available Endpoints

- `GET /` - Welcome message
- `GET /countries` - Get country suggestions (200+ countries)
- `GET /cities` - Get city suggestions (74000+ cities)
- `GET /universities` - Get university suggestions (0+ universities)
- `GET /companies` - Get company suggestions (0+ companies)
- `GET /job-titles` - Get job title suggestions (850+ job titles)
- `GET /skills` - Get skill suggestions (87000+ skills)

### Example Usage

```bash
# Get up to 10 countries containing "United"
GET http://localhost:1323/countries?q=United&limit=10

# Get up to 20 cities starting with "San"
GET http://localhost:1323/cities?q=San

# Get up to 5 universities containing "University of"
GET http://localhost:1323/universities?q=University of&limit=5
```

## Project Structure

```
autocomplete/
├── controllers/          # HTTP request handlers
│   ├── citiesController.go
│   ├── companiesController.go
│   ├── countriesController.go
│   ├── jobTitlesController.go
│   ├── skillsController.go
│   └── universitiesController.go
├── pkg/                 # Shared packages
│   ├── trie/           # Trie data structure implementation
│   │   └── trie.go
│   └── utils/          # Utility functions
│       └── limit.go    # Limit parsing logic
├── services/           # Business logic layer
│   └── databuilder.go  # Data initialization
├── raw/               # Raw data files
│   ├── cities.json
│   ├── companies.json
│   ├── countries.json
│   ├── jobs.json
│   ├── skills.json
│   └── universities.json
├── go.mod             # Go module definition
├── go.sum             # Go module checksums
├── server.go          # Main application entry point
└── README.md         # Project documentation
```

## Requirements

- Go 1.24.3 or higher

## Installation

1. Clone the repository:
```bash
git clone https://github.com/AliKhazaaleh/autocomplete.git
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

## Response Format

All endpoints return JSON arrays of strings:

```json
[
  "matched item 1",
  "matched item 2",
  "matched item 3"
]
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.