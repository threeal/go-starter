# Go Starter

A minimal template for building a [Go](https://go.dev/) library with an optional CLI. Ships pre-configured with formatting, linting, 100% test coverage enforcement, pre-commit hooks, and CI.

## Getting Started

Create a new repository from this template on GitHub using [this link](https://github.com/new?template_name=go-starter&template_owner=threeal), or clone it locally and point it at your own remote.

## Setup

Install [Go](https://go.dev/), [Lefthook](https://lefthook.dev/), [dprint](https://dprint.dev/), and [golangci-lint](https://golangci-lint.run/welcome/install/), then register the pre-commit hook:

```sh
lefthook install
```

## Customizing

Replace or extend the template files to fit your project:

- **`cmd/`** — Replace the placeholder CLI with your own.
- **`CLAUDE.md`** — Replace with guidance specific to your project.
- **`fibonacci.go`** — Replace the placeholder library and its tests with your own code.
- **`go.mod`** — Rename the module to your repository's import path, such as `github.com/username/project`, then update the package names to match.
- **`LICENSE`** — Replace with your preferred license, or keep the [Unlicense](https://unlicense.org/).
- **`README.md`** — Replace with a description of your project.

If your project doesn't need a CLI, remove `cmd/bonacci/`, along with its exclusion in `.testcoverage.yaml` and the `go-arg` dependency in `go.mod`.

## Development

Run the tests with:

```sh
go test ./... -cover
```

CI enforces 100% code coverage, excluding paths listed under `exclude.paths` in `.testcoverage.yaml`, which `go test` still counts as uncovered.

Each `git commit` runs the pre-commit hook registered during setup, which checks your changes and fixes what it can in place. If it fails, fix any reported issues, re-stage the changed files, and commit again.

After committing, push to `main` or open a pull request from another branch — CI will run the same checks across all files, plus additional checks of its own.

## Releasing

Update the CLI's `--version` in `cmd/bonacci/main.go`, push a version tag such as `v1.0.0`, and create a GitHub Release.
