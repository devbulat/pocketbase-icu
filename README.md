# PocketBase with SQLite ICU

## Overview

This project builds [PocketBase](https://pocketbase.io/) v0.40.6 with SQLite ICU support. SQLite's default `LIKE` is case-insensitive only for ASCII; ICU enables Unicode case folding, including Cyrillic searches. See the [SQLite ICU documentation](https://www.sqlite.org/src/doc/trunk/ext/icu/README.txt).

## Requirements

- Go 1.27.3 or newer (see `go.mod`).
- A C compiler and ICU development libraries. The `mattn/go-sqlite3` driver requires CGO.

## Build on Linux

On Ubuntu or Debian, install the system dependencies, then build from the repository root:

```sh
sudo apt-get update
sudo apt-get install -y build-essential libicu-dev
CGO_ENABLED=1 go build -buildvcs=false -tags sqlite_icu -o pocketbase .
```

## Build on macOS

On Apple Silicon or Intel, install the Xcode Command Line Tools and Homebrew dependencies:

```sh
xcode-select --install
brew install go pkg-config icu4c
```

If the Command Line Tools are already installed, skip the first command. Build from the repository root; `brew --prefix icu4c` selects the ICU path for either architecture:

```sh
CGO_ENABLED=1 \
CGO_CFLAGS="-I$(brew --prefix icu4c)/include" \
CGO_LDFLAGS="-L$(brew --prefix icu4c)/lib" \
PKG_CONFIG_PATH="$(brew --prefix icu4c)/lib/pkgconfig" \
go build -buildvcs=false -tags sqlite_icu -o pocketbase .
```

## Verify Unicode Search

Start PocketBase:

```sh
./pocketbase serve
```

Open the [Admin UI](http://127.0.0.1:8090/_/), navigate to **SQL Console**, and execute:

```sql
SELECT 'ПрИвЕт' LIKE '%привет%' AS matches;
```

The result should be `1`, confirming that case-insensitive Unicode search is working in PocketBase.

## Important Notes

- CGO must remain enabled, and ICU libraries must be available when building and running the binary.
- Binaries are platform-specific; build separately for each operating system and architecture.
- Do not use PocketBase's built-in `update` command: it replaces this custom binary with an upstream build. Rebuild from this source instead.
- This repository distributes source and build instructions only, not precompiled binaries.

## License

This project uses [PocketBase](https://github.com/pocketbase/pocketbase), licensed under the [MIT License](https://github.com/pocketbase/pocketbase/blob/v0.40.6/LICENSE.md). PocketBase copyright © 2022–present Gani Georgiev. Retain the upstream copyright and license notice when distributing copies or substantial portions of its code.
