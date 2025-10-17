# JobAgg (backend)

Small Go backend that proxies JSearch (RapidAPI) and exposes a single HTTP endpoint to search jobs.

Frontend: https://github.com/Passenger7390/JobAgg_Frontend

- Main entry: [`main.main`](main.go) — [main.go](main.go)
- JSearch client: [`jsearch.NewClient`](internal/jsearch/client.go), [`jsearch.Client.SearchJob`](internal/jsearch/client.go) — [internal/jsearch/client.go](internal/jsearch/client.go)
- Request/response models: [`models.SearchJobParams`](models/jsearch.go), [`models.JSearchResponse`](models/jsearch.go) — [models/jsearch.go](models/jsearch.go)
- Module file: [go.mod](go.mod)
- Environment template: [.env](.env)

## Requirements

- Go (use version compatible with the project module in [go.mod](go.mod))
- RapidAPI JSearch credentials (set in environment)

## Environment

Create a `.env` file (or export env vars) with:

- JSEARCH_API — your RapidAPI key
- JSEARCH_API_HOST — the RapidAPI host header value (as used in [main.go](main.go))

Example `.env`:

JSEARCH_API=your_api_key_here
JSEARCH_API_HOST=jsearch.p.rapidapi.com

## Build & Run

From repository root:

- Run directly:

```bash
go run main.go
```

- Build binary:

```bash
go build -o JobAgg main.go
./JobAgg
```

The app uses vendor dependencies included in `vendor/` by default.

## API

POST /jobs

- Required parameter: query
- Body: JSON matching [`models.SearchJobParams`](models/jsearch.go)
- Response: JSON matching [`models.JSearchResponse`](models/jsearch.go)

Example request:

```json
{
  "query": "software engineer",
  "page": 1,
  "num_pages": 1,
  "country": "ph",
  "language": "en",
  "date_posted": "all"
}
```
