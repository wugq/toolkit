# Architecture Notes

## cmd vs runner

- **`cmd/`** — handles user input and output. Each file defines a Cobra command: parses flags, validates input, and prints results to stdout.
- **`runner/`** — contains the core logic for each command. Runner packages are pure logic with no CLI concerns; they are called by the corresponding cmd file.

## Testing

- Unit tests live in each runner package alongside the source (e.g. `runner/dateRunner/date_test.go`).
- Test files use the same package name (white-box), giving access to unexported helpers.
- `ipRunner` has no unit tests — `LocalIPs` depends on the host network state and `PublicIP` calls an external API; treat these as integration tests.
- `stdinUtil` has only a smoke test because its functions read directly from `os.Stdin`.
- Run all tests: `go test ./...`
