# leasewebgo

A Golang client for the Leaseweb API.

## Installation

To import this library into your Go project:

```go
import "github.com/rid/leasewebgo"
```

**Note:** A minimum of Go 1.14 is required for development.

Download module  with:

```sh
go get github.com/rid/leasewebgo
```

## Stability and Compatibility

Leasewebgo is currently provided with a major version of [v0](https://blog.golang.org/v2-go-modules). We'll try to avoid breaking changes to this library, but they will certainly happen as we work towards a stable v1 library. See [CHANGELOG.md](CHANGELOG.md) for details on the latest additions, removals, fixes, and breaking changes.

While leasewebgo provides an interface to most of the [LeaseWeb API](https://developer.leaseweb.com/), the API is regularly adding new features. To request or contribute support for more API end-points or added fields, [create an issue](https://github.com/Rid/leasewebgo/issues/new).

## Usage

To authenticate to the Equinix Metal API, you must have your API token exported in env var `LEASEWEB_AUTH_TOKEN`.

This code snippet initializes Equinix Metal API client, and lists your Projects:

```go
package main

import (
	"log"

	"github.com/rid/leasewebgo"
)

func main() {
	c, err := leasewebgo.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	bms, _, err := c.bareMetals.List(nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, bm := range bms {
		log.Println(bm.ip, bm.macAddress)
	}
}

```

## Contributing

See [CONTIBUTING.md](CONTRIBUTING.md).
