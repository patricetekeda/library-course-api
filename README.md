# library-course-api

A small Go-based API project ("library-course-api"). This repository contains the source for a library/course API service.

## Module

Module path from `go.mod`:

```
module github.com/patricetekeda/library-course-api
```

Go version (from `go.mod`): `go 1.25.3`

## Prerequisites

-   Go 1.25.3 or newer installed (https://go.dev/dl/)
-   (Optional) `make` for convenience

## Quick start

Clone the repository and build:

```bash
# from the parent directory
git clone https://github.com/patricetekeda/library-course-api.git
cd library-course-api

# build
go build ./...
```

Run tests:

```bash
go test ./... -v
```

Run (if this repo provides a runnable main):

```bash
# if there's a main package at ./cmd/server or similar
go run ./cmd/server
```

## Project layout (suggested)

-   `cmd/` - entrypoints for applications
-   `internal/` - private application code
-   `pkg/` - library code intended for external use
-   `api/` - API schema or docs
-   `test/` - test helpers and fixtures

Adjust as needed to match this repository's structure.

## Development notes

-   Keep `go.mod` and `go.sum` under version control.
-   Use `gofmt`/`go vet` and consider `golangci-lint` for linting.
-   Add `.env` to `.gitignore` for local secrets.

## Contributing

Contributions are welcome — open an issue or submit a pull request with tests and a short description of changes.

## License

Add a `LICENSE` file to this repository with your chosen license (MIT, Apache-2.0, etc.).
