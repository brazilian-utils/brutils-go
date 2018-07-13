# Brazilian Utils for Go

[![CircleCI](https://circleci.com/gh/brazilian-utils/brutils-go/tree/master.svg?style=svg)](https://circleci.com/gh/brazilian-utils/brutils-go/tree/master)

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
