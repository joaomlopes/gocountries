# Gocountry

This is a [Go](https://golang.org) wrapper library around the API provided by
[Restcountries](https://restcountries.eu).

## Installation

Just go with

```shell
go get github.com/joaomlopes/gocountries
```

## Example Usage

```go
package main

import (
    "fmt"
    "github.com/joaomlopes/gocountries"
)

func main() {
    countries, err := gocountries.CountriesByName("italy")

    if err == nil {
        country := (countries)[0]
        fmt.Println(fmt.Sprintf("The capital of Italy is %s", country.Capital))
    }
}

```

## Contribution

Please, if you want to contribute go to the main project which I forked:
[https://github.com/alediaferia/gocountries](https://github.com/alediaferia/gocountries)

## License

This library is provided with a MIT License.
