# Architecture Notes

## cmd vs runner

- **`cmd/`** — handles user input and output. Each file defines a Cobra command: parses flags, validates input, and prints results to stdout.
- **`runner/`** — contains the core logic for each command. Runner packages are pure logic with no CLI concerns; they are called by the corresponding cmd file.
