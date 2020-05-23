# YAML matrix

## About

Converts YAML in a such way - from:
```yaml
task:
  name: first
  matrix:
    - sub: Lint
    - sub: Test
```
to:
```yaml
- task:
    name: first
    sub: Lint
- task:
    name: first
    sub: Test
   ```
Nesting matrix inside other matrix allowed. Matrix modifier must be always array of mappings.

## Instructions

### Preparations

Install `task` https://taskfile.dev/#/installation 

```sh
task --list # for available tasks
```

### build CLI

```sh
task build # builds binary
./cmd/cli/cli.bin <filename> # apply conversion
```

## Development 

Install golangci-lint https://golangci-lint.run/usage/install/ v1.26.0

```sh
task attach_hooks # attaches git hooks for development
task test # run test
task lint # run linters
```