# go-zendesk
[![Gitter](https://badges.gitter.im/terraform-provider-zendesk/Lobby.svg)](https://gitter.im/terraform-provider-zendesk/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![Build Status](https://travis-ci.org/nukosuke/go-zendesk.svg?branch=master)](https://travis-ci.org/nukosuke/go-zendesk)
[![Build status](https://ci.appveyor.com/api/projects/status/ce4p1mswjkdftv6o/branch/master?svg=true)](https://ci.appveyor.com/project/nukosuke/go-zendesk/branch/master)
[![Coverage Status](https://coveralls.io/repos/github/nukosuke/go-zendesk/badge.svg?branch=master)](https://coveralls.io/github/nukosuke/go-zendesk?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/nukosuke/go-zendesk)](https://goreportcard.com/report/github.com/nukosuke/go-zendesk)
[![GoDoc](https://godoc.org/github.com/zenform/go-zendesk?status.svg)](https://godoc.org/github.com/zenform/go-zendesk)

Zendesk API client library for Go

- [Reference](https://godoc.org/github.com/nukosuke/go-zendesk)

## Installation

``` shell
$ go get github.com/nukosuke/go-zendesk
```

## Usage

```go
package main

import (
    "github.com/nukosuke/go-zendesk/zendesk"
)

func main() {
    // You can set custom *http.Client here
    client := zendesk.NewClient(nil)

    // example.zendesk.com
    client.SetSubdomain("example")

    // Authenticate with API token
    client.SetCredential(zendesk.NewAPITokenCredential("john.doe@example.com", "apitoken"))

    // Authenticate with agent password
    client.SetCredential(zendesk.NewBasicAuthCredential("john.doe@example.com", "password"))

    // Create resource
    client.CreateGroup(zendesk.Group{
        Name: "support team",
    })
}
```

## License

MIT License.

See the file [LICENSE](./LICENSE).
