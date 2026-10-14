# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

> **Template notice:** This file describes the template repository itself. In a project derived from this template, tell the user it should be replaced with project-specific guidance.

## About This Repository

A minimal Go library and CLI starter template. Its only application code is a placeholder Fibonacci implementation in `fibonacci.go` and `cmd/bonacci/` — real logic is added by projects derived from it.

## Gotchas

- `lefthook run pre-commit` skips every job when nothing is staged, even ones that ignore the staged file list — pass `--all-files` to run it outside an actual commit.
- The pre-commit hook fixes files in place, and the run fails if any file changes. Report that failure and leave the fixes for the user to review and re-stage before committing again.
- `go test` doesn't apply the 100% coverage threshold — `go tool go-test-coverage` does, on the profile `go test` writes; both commands are in the CI workflow. `go test`'s own percentage reads lower because it still counts code marked `// coverage-ignore`.
- `.gitignore` ignores every dotfile and dotdir (`.*`) — a new one needs an explicit `!` exception there or it silently won't be tracked.
