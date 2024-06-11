<h1 align="center">ollamit</h1>

<p align="center">
 <a href="https://github.com/koki-develop/ollamit/releases/latest"><img src="https://img.shields.io/github/v/release/koki-develop/ollamit" alt="GitHub release (latest by date)"></a>
 <a href="https://github.com/koki-develop/ollamit/releases/latest"><img alt="GitHub all releases" src="https://img.shields.io/github/downloads/koki-develop/ollamit/total?style=flat"></a>
 <a href="https://github.com/koki-develop/ollamit/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/koki-develop/ollamit/ci.yml?logo=github" alt="GitHub Workflow Status"></a>
 <a href="https://goreportcard.com/report/github.com/koki-develop/ollamit"><img src="https://goreportcard.com/badge/github.com/koki-develop/ollamit" alt="Go Report Card"></a>
 <a href="./LICENSE"><img src="https://img.shields.io/github/license/koki-develop/ollamit" alt="LICENSE"></a>
</p>

<p align="center">
A command-line tool to generate commit messages with ollama.
</p>

## Contents

- [Installation](#installation)
  - [Homebrew Tap](#homebrew-tap)
  - [`go install`](#go-install)
  - [Releases](#releases)
- [Usage](#usage)
- [LICENSE](#license)

## Installation

### Homebrew Tap

```console
$ brew install koki-develop/tap/ollamit
```

### `go install`

```console
$ go install github.com/koki-develop/ollamit@latest
```

### Releases

Download the binary from the [releases page](https://github.com/koki-develop/ollamit/releases/latest).

## Usage

```console
$ ollamit
A command-line tool to generate commit messages with ollama.

Usage:
  ollamit [flags]

Flags:
      --dry-run        dry run
  -h, --help           help for ollamit
  -m, --model string   model name
  -v, --version        version for ollamit
```

## LICENSE

[MIT](./LICENSE)
