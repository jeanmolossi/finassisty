# Guidelines for FinAssisty

These instructions apply to all files **outside of `app/`**. The main backend is written in Go.

## Go code style

- Format all Go code using `gofmt` before committing.
- Organize imports in three groups: standard library, third-party, and local packages.
- Follow `golangci-lint` rules defined in `.golangci.yml` by running `golangci-lint run ./...`.
- Run `go test ./...` to ensure tests pass.
- After adding dependencies run `go mod tidy` so `go.mod` and `go.sum` stay in sync.
- Keep environment examples in `.env.example` up to date.

## Documentation

- When editing files in `docs/architecture/`, regenerate PNG diagrams with `make docs`.

## Pull Requests

- Use commit messages in the form `feat:`, `fix:` or similar prefixes.
- Provide a short summary of changes and mention test results or why tests could not be run.

## Getting started

- Run `make install-deps` and ensure all packages are installed.

## Developing

- Always document the work. Ensure related functions or components include meaningful docstrings when relevant.
  - If more context about the product is needed, put a doc file in `docs/product/`

## Ending development

- Run `make lint` and make sure there are no lint issues.
  - If two or more lint rules conflict, disable the least relevant rule.
