# YAML matrix

## Instructions

Install `task` https://taskfile.dev/#/installation

```sh
task --list # for available tasks
```

## CLI

```sh
task build # builds binary
./cmd/cli/cli.bin <filename> # apply conversion
```

## Development 

Install golangci-lint https://golangci-lint.run/usage/install/ v1.26.0

```sh
task attach_hooks # attaches git hooks for develompent
task test # run test
task lint # run linters
```