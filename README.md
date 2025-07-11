# sqlnull

[![Go version](https://img.shields.io/github/go-mod/go-version/nktmys/sqlnull)](https://github.com/nktmys/sqlnull/blob/main/go.mod) 
[![GoDoc](https://pkg.go.dev/badge/github.com/nktmys/sqlnull)](https://pkg.go.dev/github.com/nktmys/sqlnull) 
[![GitHub release](https://img.shields.io/github/release/nktmys/sqlnull)](https://github.com/nktmys/sqlnull/releases) 
[![Go Report Card](https://goreportcard.com/badge/github.com/nktmys/sqlnull)](https://goreportcard.com/report/github.com/nktmys/sqlnull)
[![Build Status](https://img.shields.io/github/actions/workflow/status/nktmys/sqlnull/test.yaml?label=go%20test)](https://github.com/nktmys/sqlnull/actions/workflows/test.yaml)
[![Coverage Status](https://coveralls.io/repos/github/nktmys/sqlnull/badge.svg?branch=main)](https://coveralls.io/github/nktmys/sqlnull?branch=main)

`sqlnull` is utilities for working with sql.Null types in Go.

## Installation
1. Get the package
```sh
go get github.com/nktmys/sqlnull
```

2. Import the package
```go
import "github.com/nktmys/sqlnull"
```

## Examples

This code:

```go
str := "hello"
v1 := sqlnull.From(str)
v2 := sqlnull.FromStr(&str)

p1 := sqlnull.Prt(str)
p2 := sqlnull.PrtOrNil(&str)
```

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License
This project is licensed under the MIT License. See the [LICENSE](https://github.com/nktmys/sqlnull/blob/main/LICENSE) file for details.
