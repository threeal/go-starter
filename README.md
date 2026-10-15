# Go Starter

A minimal template for building a [Go](https://go.dev/) library with an optional CLI. Ships pre-configured with formatting, 100% test coverage enforcement, pre-commit hooks, and CI.

## Getting Started

Create a new repository from this template on GitHub using [this link](https://github.com/new?template_name=go-starter&template_owner=threeal), or clone it locally and point it at your own remote.

## Setup

Install [Go](https://go.dev/), then install dependencies:

```sh
go mod download
```

Install [Lefthook](https://lefthook.dev/) and [dprint](https://dprint.dev/), then register the pre-commit hook:

```sh
lefthook install
```

## Customizing

Replace or extend the template files to fit your project:

- **`cmd/bonacci/`** — Replace or remove the placeholder CLI, renaming the directory to your command's name. Remove this directory and the `go-arg` dependency if your project doesn't need a CLI.
- **`CLAUDE.md`** — Replace with guidance specific to your project.
- **`fibonacci.go`** — Replace with your own library logic.
- **`go.mod`** — Update the module path to your repository's import path, such as `github.com/username/project`, and update the CLI's import and each library file's `package` clause to match.
- **`LICENSE`** — Replace with your preferred license, or keep the [Unlicense](https://unlicense.org/).
- **`README.md`** — Replace with a description of your project.

## Development

Write code at the module root. Test files live alongside source as `*_test.go`. Run the test suite with:

```sh
go test ./... -coverprofile=.cover.out -coverpkg=./...
go tool -modfile=tools/go.mod go-test-coverage -p .cover.out -f 100 -k 100 -t 100
```

The project enforces 100% code coverage on every run.

Each `git commit` runs the pre-commit hook registered during setup, which checks your changes and fixes what it can in place. If it fails, fix any reported issues, re-stage the changed files, and commit again.

After committing, push to `main` or open a pull request from another branch — CI will run the same checks across all files, plus additional checks of its own.

## Releasing

Update the CLI's `--version` in `cmd/bonacci/main.go`, push a version tag such as `v1.0.0`, and create a GitHub Release. The module is then available to other projects through `go get`, and its documentation appears on [pkg.go.dev](https://pkg.go.dev/).
