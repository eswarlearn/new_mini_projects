# filerunner

A tiny Cobra-based CLI that reads a file and executes it in your terminal.

## Setup

```bash
# 1. Unpack / place the files so you have this layout:
#    filerunner/
#      go.mod
#      main.go
#      cmd/root.go
#      cmd/run.go

cd filerunner

# 2. Download the cobra dependency and lock go.sum
go mod tidy

# 3. Build a binary (optional — you can also just `go run .`)
go build -o filerunner .
```

## Usage

Two modes, controlled by the `--script` flag.

### 1. Line-by-line mode (default)

Runs each non-empty, non-comment line of the file as its own shell command,
streaming output live. Good for a plain list of commands:

```
# commands.txt
echo "starting"
ls -la
go version
```

```bash
./filerunner run commands.txt
```

Useful flags:
- `-n, --dry-run` — print what would run without executing it
- `--stop-on-error=false` — keep going even if a line fails (default: stop)

### 2. Script mode

Hands the whole file to an interpreter in one shot (like running `bash file.sh`).
Use this for real shell scripts with `if`, loops, functions, etc.

```bash
./filerunner run deploy.sh --script
./filerunner run setup.py --script --interpreter python3
```

Flags:
- `-s, --script` — treat the file as one script instead of line-by-line
- `-i, --interpreter` — interpreter to use with `--script` (default: `bash`)

### Examples

```bash
# Dry run to preview what would happen
./filerunner run commands.txt --dry-run

# Run a real bash script, letting it fail without stopping the flow
./filerunner run backup.sh --script

# Or without building a binary first
go run . run commands.txt
```

## How it works

- `cmd/root.go` sets up the base `filerunner` Cobra command.
- `cmd/run.go` implements the `run` subcommand:
  - **line-by-line**: reads the file with `bufio.Scanner`, skips blank
    lines and `#` comments, and runs each remaining line via
    `sh -c "<line>"` so pipes/redirects/env vars work as expected.
  - **script mode**: just execs `<interpreter> <file>` directly
    (e.g. `bash file.sh`).
- In both cases, the child process's `Stdout`/`Stderr`/`Stdin` are wired
  directly to the current terminal, so output streams live instead of
  being buffered and dumped at the end.

## Extending it

Some natural next steps if you want to keep building on this:
- Add a `--timeout` flag using `context.WithTimeout` around `exec.CommandContext`.
- Add `--env KEY=VALUE` flags to inject extra environment variables into the child process.
- Add a `--parallel` flag to run line-by-line commands concurrently with goroutines + a `sync.WaitGroup`.
