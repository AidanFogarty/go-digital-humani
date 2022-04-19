# Go Digital Humani 🌲

[![Build Status](https://github.com/AidanFogarty/go-digital-humani/actions/workflows/pipeline.yml/badge.svg)](https://github.com/AidanFogarty/go-digital-humani)
[![GoDoc](https://godoc.org/github.com/AidanFogarty/go-digital-humani?status.svg)](https://pkg.go.dev/github.com/AidanFogarty/go-digital-humani)
[![Go Report Card](https://goreportcard.com/badge/github.com/AidanFogarty/go-digital-humani)](https://goreportcard.com/report/github.com/AidanFogarty/go-digital-humani)

Unoffical Go SDK for DigitalHumani's Reforestation as a Service API. A platform which easily integrates with your products to reforest the planet.

## Prerequisites

A DigitalHumani Account is needed to retrieve your:
- Enterprise ID
- Api Key

## Installation

To install, simply run `go get` in your project.

```
go get github.com/AidanFogarty/go-digital-humani
```

To update the SDK, use `go get -u` to retrieve the latest version.

```
go get -u github.com/AidanFogarty/go-digital-humani
```

## Getting Started

### Plant a tree

To plant a single tree, add the following to your `main.go`:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AidanFogarty/go-digital-humani/pkg/digitalhumani"
)

func main() {
	dh := digitalhumani.New("<your-api-key>", "<your-enterprise-id>", "sandbox")

	plantTreeResp, err := dh.PlantTree(context.TODO(), "93322350", "MyUser", 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(plantTreeResp)
}
```

Compile & Execute:
```bash
> go run main.go
&{3ea84180-bffe-11ec-9d64-0242ac120002 2022-04-19T16:29:54.171Z 1 abcdefg 93322350 MyUser}
```

For other features, see the [docs](https://pkg.go.dev/github.com/AidanFogarty/go-digital-humani)

## Contributing