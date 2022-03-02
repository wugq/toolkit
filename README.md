# toolkit

This toolkit is built with [Go](https://go.dev/) and [Cobra](https://cobra.dev/).

## Dev

### How to add subcommand:

`cobra add cmdname`

### How to build binary:

`go build`

## Usage

### Available subcommands

#### touch

Update timestamp of a file or a folder, like touch in Linux

#### generate

* Generate a password `generate password`
* Generate a UUID string `generate uuid`

#### date

Show date as well as unix time, with simple calculation.

#### dig

Show DNS lookup info for a domain