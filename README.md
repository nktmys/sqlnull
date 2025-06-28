# sqlnull

[![GitHub release](https://img.shields.io/github/release/nktmys/sqlnull.svg)](https://github.com/nktmys/sqlnull/releases) [![GoDoc](https://godoc.org/github.com/nktmys/sqlnull?status.svg)](https://godoc.org/github.com/nktmys/sqlnull) [![Go Report Card](https://goreportcard.com/badge/github.com/nktmys/sqlnull)](https://goreportcard.com/report/github.com/nktmys/sqlnull)

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
