# TODO

## Installation

```shell
go get -u -v github.com/brazilian-utils/brutils-go
```

## Usage


```golang
package main

import "github.com/brazilian-utils/brutils-go/cpf" // CPF package

func main() {
    if cpf.IsValid("40364478829") {
        // valid CPF. This one was generated randomly, but is a valid one.
    }
}
```
