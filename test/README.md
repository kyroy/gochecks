# GitHub Action for `go test`


## Usage

```workflow
action "go test" {
    uses = "kyroy/gochecks/test@master"
}
```

It is possible to specify the following `args`:
- `--input`, `-i` path to test log file executed with `-json` flag (e.g. `go test -json`)
- `--dir` path to the test directory
- `--fail` if set, the command will fail on failed tests

It is possible to use the following subcommands:

### `pr`

TBD
