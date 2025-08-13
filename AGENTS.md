# Repository Guidelines

## Project Structure & Module Organization
- `cmd/web`: Go entrypoint and HTTP handlers.
- `ui/html`: Templates — `base.tmpl` plus page templates in `ui/html/pages`.
- `ui/static`: Public assets served under `/static` (CSS, images, carousel).
- `Buildfile`/`Procfile`: Build/run commands for deployment (Elastic Beanstalk compatible).
- Module: `cakecountry.ca` (Go 1.22+).

## Build, Test, and Development Commands
- `go run ./cmd/web`: Run the server locally (uses `PORT`, defaults to 5000).
- `go build -o bin/application ./cmd/web`: Build binary locally.
- `PORT=5000 ./bin/application`: Start built binary with explicit port.
- `go test ./...`: Run tests when present.
- `go vet ./...`: Static checks for common issues.

## Coding Style & Naming Conventions
- Formatting: `gofmt -s -w .` before pushing; keep imports grouped.
- Linting: Prefer idiomatic Go; avoid stutter in names (e.g., `web.Handler`).
- Packages/dirs: lowercase, no underscores; files `snake_case.go` is fine, prefer concise.
- Handlers: one function per route in `cmd/web` (e.g., `home`, `pricing`).
- Templates: place page-specific files in `ui/html/pages`; render via the `base` layout.

## Testing Guidelines
- Framework: standard `testing` package.
- File names: `*_test.go`; test funcs `TestXxx`.
- Scope: table tests for handlers/utilities; validate status codes and rendered content.
- Run: `go test ./...`; aim to cover new logic paths.

## Commit & Pull Request Guidelines
- Commits: present tense, concise (e.g., "Add pricing template parser").
- Reference issues: `Fixes #123` when applicable.
- PRs: include summary, rationale, screenshots for UI/template changes, and steps to verify.
- Keep diffs focused; note config changes (e.g., `Buildfile`, `Procfile`, env vars).

## Security & Configuration Tips
- Never commit secrets; configure env in Elastic Beanstalk. Server reads `PORT`.
- Serve only public assets from `ui/static`; avoid user-uploaded files.
- Validate template paths and keep handler input parsing minimal.

## Architecture Overview
- HTTP server using `http.ServeMux`; static assets under `/static`.
- Handlers render templates with `text/template` and minimal data structs.
- Carousel images auto-discovered from `ui/static/images/carousel`.
