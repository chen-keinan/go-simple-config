[![Go Report Card](https://goreportcard.com/badge/github.com/chen-keinan/go-simple-config)](https://goreportcard.com/report/github.com/chen-keinan/go-simple-config)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/chen-keinan/go-simple-config/blob/master/LICENSE)
<img src="./pkg/img/coverage_badge.png" alt="test coverage badge">
[![Gitter](https://badges.gitter.im/beacon-sec/community.svg)](https://gitter.im/beacon-sec/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

<br><img src="./pkg/img/simple-config.png" width="300" alt="simple-config.png logo"><br>

# go-simple-config

Go Simple config is an open source configuration lib for storing and accessing configuration data with minimal
dependencies

* [Installation](#installation)
* [Supported Configuration Files](#supported-configuration-files)
* [Usage](#usage)

## Installation

```shell
go get github.com/chen-keinan/go-simple-config
```

## Supported configuration files:

- yaml
- json
- properties
- environment variables
- ini

## Usage

### json config example:

```json
{
  "SERVER": {
    "host": "127.0.0.1",
    "port": "8080"
  }
```

### yaml config example:

```yaml
---
SERVER:
  host: 127.0.0.1
  port: '8080'
```

### properties config example:

```json
SERVER.host=127.0.0.1
SERVER.port=8080
```

### env. variable config example:

```shell
export SERVER_HOST="127.0.0.1"
export SERVER_PORT="8080"
```

### Full example :

```go
package main

import (
	"fmt"
	"github.com/chen-keinan/go-simple-config/simple"
	"os"
)

func main() {
	c := simple.New()
	err := c.Load("config.json")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fmt.Print(c.GetStringValue("SERVER.host"))
}
```
