# bashenv-server

A tiny Cobra + Gin demo: the `serve` command sources a bash file
(default `~/.bashrc`) in a real bash subshell, reads whatever it
exports, and boots a Gin HTTP server with those values.

## Recognized variables

Add these to `~/.bashrc` (or any file you point `--file` at) and
`export` them:

```bash
export APP_NAME="my-cool-api"
export HOST="0.0.0.0"
export PORT="9090"
export GIN_MODE="debug"   # debug | release | test
```

Anything you don't export falls back to a default
(`HOST=0.0.0.0`, `PORT=8080`, `GIN_MODE=debug`, `APP_NAME=bashenv-server`).

## Setup

```bash
cd bashenv-server
go mod tidy      # fetches gin and cobra, writes go.sum
go build -o bashenv-server .
```

## Run

```bash
# uses ~/.bashrc by default
./bashenv-server serve

# point at a different file
./bashenv-server serve --file ~/.bash_profile

# CLI flag overrides whatever PORT was in the bash file
./bashenv-server serve --port 3000

# skip the startup summary print
./bashenv-server serve -q
```

Then hit it:

```bash
curl localhost:8080/health
curl localhost:8080/config
curl localhost:8080/
```

## Why shell out to bash instead of parsing the file?

`.bashrc` is a shell script, not a `KEY=VALUE` file — it can have
conditionals, sourced sub-files, command substitution, PATH
manipulation, etc. Hand-parsing it with regex would break the moment
someone's bashrc did anything non-trivial. Instead we run
`bash -i -c "source <file> && env"` and diff/parse the resulting
environment, which is what bash itself would produce for a real
interactive shell.

## Project layout

```
bashenv-server/
├── go.mod
├── main.go
├── cmd/
│   ├── root.go     # cobra root command
│   └── serve.go    # cobra "serve" subcommand -> starts gin
├── internal/
│   └── config/
│       └── bashenv.go   # sources the bash file, parses env, builds config
└── README.md
```

## Extending it

- Add more recognized keys in `BuildServerConfig` (internal/config/bashenv.go).
- Add a `--env-only` flag that just prints the parsed vars and exits,
  useful for debugging what a given bashrc actually exports.
- Swap `env` parsing for `os.Environ()` diffing (snapshot env before
  and after sourcing) if you want to detect *only* what the file
  added, rather than the whole shell environment.
