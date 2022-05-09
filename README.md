# Go utilities

![Tests](https://github.com/murat/dicterm/actions/workflows/test.yml/badge.svg)

This package includes various utilities and extensions for your Go code.

Inspired by [lodash](https://lodash.com)

## Install

```shell
go get github.com/murat/go-utils@master
```

## Usage

### String extensions

```go
package main

import (
	"fmt"

	"github.com/murat/go-utils/strings"
)

func main() {
	// Capitalize
	fmt.Println(strings.Capitalize("hello world")) // Hello World
	// CamelCase
	fmt.Println(strings.ToCamelCase("hello world")) // helloWorld
	// SnakeCase
	fmt.Println(strings.ToSnakeCase("hello world")) // hello_world
	// Deburr
	fmt.Println(strings.Deburr("éàèù")) // eaeu
	// KebabCase
	fmt.Println(strings.ToKebabCase("hello world")) // hello-world
	// Paddings
	fmt.Println("|" + strings.Pad("hello", 20) + "|")   // |       hello        |
	fmt.Println(strings.PadWith("hello", 20, "-"))      // |-------hello--------|
	fmt.Println(strings.PadLeft("hello", 20))           // |               hello|
	fmt.Println(strings.PadLeftWith("hello", 20, "-"))  // |---------------hello|
	fmt.Println(strings.PadRight("hello", 20))          // |hello               |
	fmt.Println(strings.PadRightWith("hello", 20, "-")) // |hello---------------|
}
```

All contributors are welcome to contribute to this project.

Cheers :beer:
