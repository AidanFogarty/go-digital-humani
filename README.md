# Go DigitalHumani 🌲
[![Build Status](https://github.com/AidanFogarty/go-digital-humani/actions/workflows/pipeline.yml/badge.svg)](https://github.com/AidanFogarty/go-digital-humani)
[![GoDoc](https://godoc.org/github.com/AidanFogarty/go-digital-humani?status.svg)](https://pkg.go.dev/github.com/AidanFogarty/go-digital-humani)

Unoffical Go SDK for DigitalHumani's Reforestation as a Service API. A platform which easily integrates with your products to reforest the planet.

## Prerequisites
---

A DigitalHumani Account is needed to retrieve your:
- Enterprise ID
- Api Key

For information on how to setup and get your DigitalHumani account details, see the [API Docs](https://docs.digitalhumani.com/)

## Installation
---

To install, simply run `go get` in your project.

```
go get github.com/AidanFogarty/go-digital-humani
```

To update the SDK, use `go get -u` to retrieve the latest version.

```
go get -u github.com/AidanFogarty/go-digital-humani
```

## Getting Started
---

To get started with the DigitalHumani sdk, simply create a new instance.

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AidanFogarty/go-digital-humani/pkg/digitalhumani"
)

func main() {
	dh := digitalhumani.New("<your-api-key>", "<your-enterprise-id>", "sandbox") // Can be either 'sandbox' or 'production'

	// Plant some trees......... 🌲🌲🌲
}
```

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

### Retrieve Details of a Plant Tree Request

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

	plantTreeResp, err := dh.GetTree(context.TODO(), "<uuid-of-request>")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(plantTreeResp)
}
```

### Get number of trees planted by a user

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

	plantTreeCount, err := dh.GetTreeCount(context.TODO(), "<user-id>")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(plantTreeCount)
}
```

### Get Enterprise Details

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

	enterprise, err := dh.GetEnterprise(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(enterprise)
}
```

### Get number of trees planted by an Enterprise between two dates

Can use the the `NewDateRange` function to generate the date range parameter.

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

	enterpriseTreeCount, err := dh.GetEnterpriseTreeCount(context.TODO(), digitalhumani.NewDateRange("2022-01-01", "2022-12-31"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(enterpriseTreeCount)
}
```

### Get number of trees planted by an Enterprise for a given month

Month must be in format `YYYY-MM`

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

	enterpriseTreeCount, err := dh.GetEnterpriseMonthTreeCount(context.TODO(), "2022-04")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(enterpriseTreeCount)
}
```

### Get list of all reforestation projects

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

	projects, err := dh.GetAllProjects(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(projects)
}
```

### Get a single project details by ID

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

	project, err := dh.GetProject(context.TODO(), "32146688")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(project)
}
```

For other features and detailed description, see the [docs](https://pkg.go.dev/github.com/AidanFogarty/go-digital-humani)

## Contribution
---

Contributions are always welcome! Please read the [contribution guide](CONTRIBUTING.md).


If this project has helped you, consider leaving a ⭐!